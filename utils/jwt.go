package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/model"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

func (j *JWT) CreateToken(baseClaims model.BaseClaims) (string, error) {
	refreshAt := time.Now().Add(time.Hour * time.Duration(global.Config.JWT.RefreshTime))
	expireAt := time.Now().Add(time.Hour * time.Duration(global.Config.JWT.ExpiresTime))
	claims := model.CustomClaims{
		BaseClaims: baseClaims,
		RefreshAt:  refreshAt,
		RegisteredClaims: jwt.RegisteredClaims{
			// ID
			ID: fmt.Sprint(baseClaims.ID),
			// 主题
			Subject: baseClaims.Username,
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(expireAt),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 签发者
			Issuer: global.Config.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token is malformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token is validation failed")
			} else {
				return nil, errors.New("token is invalid")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token is invalid")

	} else {
		return nil, errors.New("token is invalid")
	}
}
