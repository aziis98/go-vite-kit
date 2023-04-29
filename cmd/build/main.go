package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

var logger = log.New(os.Stderr, "", 0)

func run(command string) {
	logger.Printf(`> %s`, color.HiWhiteString(command))

	cmd := exec.Command("sh", "-c", command)
	cmdOutReader, _ := cmd.StdoutPipe()
	cmdErrReader, _ := cmd.StderrPipe()

	s := bufio.NewScanner(io.MultiReader(cmdOutReader, cmdErrReader))
	if err := cmd.Start(); err != nil {
		logger.Fatal(err)
	}

	for s.Scan() {
		logger.Printf("  %s", color.WhiteString(s.Text()))
	}
	if err := s.Err(); err != nil {
		logger.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		logger.Fatal(err)
	}

	logger.Println()
}

func main() {
	run(`npm run build`)
	run(`go build -v -o ./out/server ./cmd/server`)
}
