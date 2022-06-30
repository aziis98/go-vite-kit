package main

import (
	"bufio"
	"log"
	"os/exec"
	"server/config"
	"server/database"
	"server/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db := database.NewInMemoryDB()

	router := &routes.Router{
		Database: db,
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", "/_frontend/dist")

	app.Route("/api", router.Api)

	if strings.HasPrefix(config.Mode, "dev") {
		log.Printf(`Running dev server for frontend: "npm run dev"`)
		cmd := exec.Command("sh", "-c", "cd _frontend/ && npm run dev")
		cmdStdout, _ := cmd.StdoutPipe()

		go func() {
			s := bufio.NewScanner(cmdStdout)
			for s.Scan() {
				log.Printf("[ViteJS] %s", s.Text())
			}
			if err := s.Err(); err != nil {
				log.Fatal(err)
			}
		}()

		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Fatal(app.Listen(config.Host))
}
