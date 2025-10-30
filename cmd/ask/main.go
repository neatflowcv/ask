package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/joho/godotenv"
	"github.com/neatflowcv/ask/internal/app/flow"
	"github.com/neatflowcv/ask/internal/pkg/gemini"
)

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	return info.Main.Version
}

func main() {
	loadEnv()

	log.Println("version", version())

	prompt := strings.Join(os.Args[1:], " ")
	if prompt == "" {
		log.Fatal("usage: ask <prompt>")
	}

	inquirer := gemini.NewClient(os.Getenv("KEY"))
	service := flow.NewService(inquirer)
	ctx := context.Background()

	answer, err := service.Ask(ctx, prompt)
	if err != nil {
		log.Fatalf("ask: %v", err)
	}

	fmt.Println(answer) //nolint:forbidigo
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("load .env file: %v", err)
	}
}
