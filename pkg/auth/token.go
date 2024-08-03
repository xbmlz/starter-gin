package auth

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xbmlz/starter-gin/pkg/env"
)

var (
	jwtSecret        = env.GetString("JWT_SECRET", "secret")
	jwtExpireTime, _ = env.GetInt("JWT_EXPIRE_TIME", 3600)
)

type UserClaims struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	expiresAt := time.Now().Add(time.Duration(jwtExpireTime) * time.Second)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err := claims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	// TODO set to redis
	return token, nil
}

func ValidateToken(tokenString string) (*UserClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
