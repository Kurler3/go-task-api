package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Kurler3/go-task-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/Kurler3/go-task-api/custom_types"
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

var jwtSecret string

// Load env
func LoadEnv() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get JWT secret from environment variables
	jwtSecret = os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET environment variable not set")
		panic("JWT_SECRET environment variable not set")
	}
}

// Generate JWT token
func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"userID": user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Validate JWT token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Initialize a new instance of `Claims`
	claims := &custom_types.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

// Get userId from context
func GetUserIdFromContext(r *http.Request) uint {
	return r.Context().Value("userId").(uint)
}
