package custom_types

import "github.com/dgrijalva/jwt-go"

// Claims struct to represent the claims in the JWT token
type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}
