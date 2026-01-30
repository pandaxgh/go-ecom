package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env  string
	App  AppConfig
	HTTP HTTPConfig
	DB   DbConfig
	Auth AuthConfig
}

type AppConfig struct {
	Name string
}

type HTTPConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DbConfig struct {
	URL string
}

type AuthConfig struct {
	JWTSecret string
	JWTExpiry time.Duration
}

func Load() *Config {
	cfg := &Config{
		Env: getEnv("ENV", "dev"),
		App: AppConfig{
			Name: getEnv("APP_NAME", "ecom"),
		},
		HTTP: HTTPConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDuration("HTTP_READ_TIMEOUT", 5*time.Second),
			WriteTimeout: getDuration("HTTP_WRITE_DURATION", 5*time.Second),
		},
		DB: DbConfig{
			URL: getEnv("DB_URL", ""),
		},
		Auth: AuthConfig{
			JWTSecret: getEnv("JWT_SECRET", ""),
			JWTExpiry: getDuration("JWT_EXPIRY", 24*time.Hour),
		},
	}

	validate(cfg)
	return cfg
}

func getEnv(key, fallback string) string {
	if key := os.Getenv(key); key != "" {
		return key
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	if key := os.Getenv(key); key != "" {
		d, err := time.ParseDuration(key)
		if err != nil {
			log.Fatalf("invalid duration for %s", key)
		}
		return d
	}
	return fallback
}

func validate(cfg *Config) {
	if cfg.Env == "prod" && cfg.DB.URL == "" {
		log.Fatal("DATABASE_URL is required in prod")
	}

	if cfg.Auth.JWTSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
}
