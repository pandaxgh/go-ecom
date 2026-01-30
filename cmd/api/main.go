package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pandaxgh/ecom-backend/internal/config"
	"github.com/pandaxgh/ecom-backend/internal/database"
)

func main() {
	fmt.Println("hey there")
	cfg := config.Load()
	db := database.Connect(cfg.DB.URL)
	defer db.Pool.Close()
}
