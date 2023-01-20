package types

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
