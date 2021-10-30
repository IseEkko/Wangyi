package model

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id       uint
	UserName string
	Password string
	Local    string
	Sex      string
	Jie      string
	HeadUrl  string
	Birth    string
	jwt.StandardClaims
}
