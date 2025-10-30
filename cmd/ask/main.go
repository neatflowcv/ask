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
	"github.com/neatflowcv/ask/internal/pkg/printer/console"
	"github.com/neatflowcv/ask/internal/pkg/printer/file"
	"github.com/neatflowcv/ask/internal/pkg/printer/group"
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
		log.Panic("usage: ask <prompt>")
	}

	fmt.Println("Question:", prompt) //nolint:forbidigo

	now := time.Now()
	inquirer := gemini.NewClient(os.Getenv("KEY"))

	filePrinter, err := file.NewPrinter("answer.txt")
	if err != nil {
		log.Fatalf("new printer: %v", err)
	}
	defer filePrinter.Close()

	groupPrinter := group.NewPrinter(console.NewPrinter(), filePrinter)

	service := flow.NewService(inquirer, groupPrinter)
	ctx := context.Background()

	err = service.Ask(ctx, prompt)
	if err != nil {
		log.Panicf("ask: %v", err)
	}

	log.Println("elapsed", time.Since(now))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("load .env file: %v", err)
	}
}
