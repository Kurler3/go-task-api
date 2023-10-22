package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Kurler3/go-task-api/services"
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

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Fetch user from the database using userID
		_, err = services.GetUserById(userID)
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
