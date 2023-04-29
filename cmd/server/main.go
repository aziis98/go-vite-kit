package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/fatih/color"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"go-vite-kit/backend/config"
	"go-vite-kit/backend/database"
	"go-vite-kit/backend/routes"
)

const FrontendOutDir = "./out/frontend"

func main() {
	db := database.NewInMemoryDB()

	router := &routes.Router{
		Database: db,
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", FrontendOutDir)
	for _, entrypoint := range routes.HtmlEntrypoints {
		app.Get(entrypoint.Route, func(c *fiber.Ctx) error {
			return c.SendFile(path.Join(FrontendOutDir, entrypoint.Filename))
		})
	}

	app.Route("/api", router.Api)

	if strings.HasPrefix(config.Mode, "dev") {
		setupDevServer()
	}

	log.Fatal(app.Listen(config.Host))
}

func setupDevServer() {
	log.Printf(`Running dev server for frontend: "npm run dev"`)
	cmd := exec.Command("sh", "-c", "npm run dev")
	cmdStdout, _ := cmd.StdoutPipe()
	cmdStderr, _ := cmd.StderrPipe()

	viteLogger := log.New(os.Stderr, color.HiGreenString("[ViteJS]")+" ", log.Ltime|log.Lmsgprefix)

	go func() {
		s := bufio.NewScanner(io.MultiReader(cmdStdout, cmdStderr))
		for s.Scan() {
			viteLogger.Print(s.Text())
		}
		if err := s.Err(); err != nil {
			viteLogger.Fatal(err)
		}
	}()

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
