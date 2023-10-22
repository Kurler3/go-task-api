package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
func comparePasswords(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

// Create JWT token

// Validate JWT token
