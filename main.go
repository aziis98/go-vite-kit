package main

import (
	"log"
	"os/exec"
	"server/database"
	"server/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.NewInMemoryDB()

	router := &routes.Router{
		Database: db,
	}

	app := fiber.New()
	app.Static("/", "/_frontend/dist")

	app.Route("/api", router.Api)

	if strings.HasPrefix(Config.Mode, "dev") {
		log.Printf(`Running dev server for frontend: "npm run dev"`)

		err := exec.Command("sh", "-c", "cd _frontend/ && npm run dev").Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Fatal(app.Listen(Config.Host))
}
