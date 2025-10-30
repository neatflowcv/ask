package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"

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

	flag.Parse()

	prompt := strings.Join(flag.Args(), " ")
	if prompt == "" {
		log.Fatal("usage: ask <prompt>")
	}

	fmt.Println("Question:", prompt) //nolint:forbidigo

	now := time.Now()
	inquirer := gemini.NewClient(os.Getenv("KEY"))
	service := flow.NewService(inquirer)
	ctx := context.Background()

	answer, err := service.Ask(ctx, prompt)
	if err != nil {
		log.Fatalf("ask: %v", err)
	}

	log.Println("elapsed", time.Since(now))

	fmt.Println(answer) //nolint:forbidigo
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("load .env file: %v", err)
	}
}
