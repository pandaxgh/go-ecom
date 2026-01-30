package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pandaxgh/ecom-backend/internal/config"
)

func main() {
	fmt.Println("hey there")
	cfg := config.Load()
	fmt.Println(cfg)
}
