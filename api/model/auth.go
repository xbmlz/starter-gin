package model

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID uint `json:"id"`
	jwt.RegisteredClaims
}
