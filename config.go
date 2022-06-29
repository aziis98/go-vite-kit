package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config struct {
	Mode    string
	Host    string
	BaseURL string
}

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

	Config.Mode = loadEnv(os.Getenv("MODE"), "development")
	Config.Host = loadEnv(os.Getenv("HOST"), ":4000")
	Config.BaseURL = loadEnv(os.Getenv("HOST"), "http://localhost:4000")
}
