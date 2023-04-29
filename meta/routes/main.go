package main

import (
	"encoding/json"
	"go-vite-kit/backend/routes"
	"log"
	"os"
)

func main() {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(routes.HtmlEntrypoints); err != nil {
		log.Fatal(err)
	}
}
