package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Mode    string
	Host    string
	BaseURL string
)

func loadEnv(key string, defaultValue ...string) string {
	env := os.Getenv(key)

	if len(defaultValue) > 0 && env == "" {
		env = defaultValue[0]
	}

	log.Printf("Environment variable %s = %q", key, env)
	return env
}

func init() {
	// Setup logger
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	// Load Config
	godotenv.Load()

	Mode = loadEnv("MODE", "development")
	Host = loadEnv("HOST", ":4000")
	BaseURL = loadEnv("BASE_URL", "http://localhost:4000")
}
