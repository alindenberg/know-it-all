package usermodels

import (
	"github.com/dgrijalva/jwt-go"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type UserClaim struct {
	Username string
	Scope    string
	Audience []string
	// Standard jwt claims
	jwt.StandardClaims
}
