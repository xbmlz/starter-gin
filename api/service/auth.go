package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xbmlz/starter-gin/api/model"
	"github.com/xbmlz/starter-gin/pkg/utils/env"
)

var (
	tokenSecret    = env.GetString("TOKEN_SECRET", "secret")
	tokenExpire, _ = env.GetInt("TOKEN_EXPIRE", 3600)
)

func TokenGenerate(user *model.User) (string, error) {
	expiresAt := time.Now().Add(time.Duration(tokenExpire) * time.Second)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JWTClaims{
		UserID: uint(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err := claims.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func TokenParse(tokenString string) (*model.JWTClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
