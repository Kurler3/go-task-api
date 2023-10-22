package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Kurler3/go-task-api/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(claims)

		// Extract user ID from claims
		userID, ok := claims["userID"].(uint)

		fmt.Println(userID, ok)

		if !ok {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		// Fetch user from the database using userID
		_, err = utils.GetUserById(userID)
		if err != nil {
			http.Error(w, "Permission Denied", http.StatusUnauthorized)
			return
		}

		// Append user to the request context
		ctx := context.WithValue(r.Context(), "userId", userID)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
