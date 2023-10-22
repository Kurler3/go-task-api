package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Kurler3/go-task-api/custom_types"
	"github.com/Kurler3/go-task-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Gets params and returns the value of the key on the params passed
func VarToUint(r *http.Request, varKey string) (uint, error) {

	vars := mux.Vars(r)

	idStr := vars[varKey]

	// Convert to uint64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

// Encrypt password
func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Compare password
func ComparePasswords(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

var jwtSecret []byte

// Load env
func LoadEnv() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get JWT secret from environment variables
	jwtSecretFromEnv := os.Getenv("JWT_SECRET")
	if jwtSecretFromEnv == "" {
		fmt.Println("JWT_SECRET environment variable not set")
		panic("JWT_SECRET environment variable not set")
	}

	jwtSecret = []byte(jwtSecretFromEnv)

}

// Generate JWT token
func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"UserID": user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Validate JWT token
func ValidateToken(tokenString string) (uint, error) {
	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &custom_types.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// Check if there was an error in parsing the token
	if err != nil {
		return 0, err
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(*custom_types.Claims); ok && token.Valid {
		// Check the user ID in the claims
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("invalid token")
}

// Return json to client
func ReturnJSONToClient(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Return simple message to client
func ReturnMessageToClient(w http.ResponseWriter, message string, statusCode int) {
	// Set response content type to plain text
	w.Header().Set("Content-Type", "text/plain")
	// Write the string response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
