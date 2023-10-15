package middleware

import (
	"fmt"
	"net/http"

	"github.io/anilk/crane/database/postgres/repositories"
)

func AuthMiddleware(userRepo *repositories.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Auth Middleware")
			user, _ := userRepo.FindByID(1)
			fmt.Println(user)
		})
	}
}
