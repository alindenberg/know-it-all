package usermodels

import (
	"github.com/dgrijalva/jwt-go"
)

type UserCredentials struct {
	Username string
	Password string
	Email    string
}

type UserClaim struct {
	Username string
	jwt.StandardClaims
}
type UserSignInRequest struct {
	Username string
	Password string
}

type UserKeys struct {
	Username string
	Token    string
}
type User struct {
	UserID   string
	Password []byte
	Username string
	Email    string
}
