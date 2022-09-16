package model

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	BaseClaims
	RefreshAt time.Time
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID       uint
	Username string
	NickName string
}
