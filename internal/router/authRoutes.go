package router

import (
	"net/http"

	"github.com/pandaxgh/ecom-backend/internal/controller"
)

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /login", controller.Login)
	mux.HandleFunc("POST /register", controller.Signup)
	mux.HandleFunc("GET /me", controller.Me)
}
