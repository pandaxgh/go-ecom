package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pandaxgh/ecom-backend/internal/config"
	"github.com/pandaxgh/ecom-backend/internal/server"
)

func main() {
	fmt.Println("hey there")
	cfg := config.Load()
	// db := database.Connect(cfg.DB.URL)
	// defer db.Pool.Close()

	srv, err := server.New(cfg)

	if err != nil {
		log.Fatalf("Failed to init server: %v", err)
	}

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola")
	})
	srv.SetupHttpServer(router)

	srv.Start()

}
