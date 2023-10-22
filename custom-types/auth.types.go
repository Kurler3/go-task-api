package custom_types

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	userID string `json:"userID"`
	jwt.Claims
}
