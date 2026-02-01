package router

import "net/http"

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	RegisterAuthRoutes(mux)
	return mux
}
