package model

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id             uint
	UserName       string
	PasswordDigest string
	Type_id        int
	jwt.StandardClaims
}
