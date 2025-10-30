package main

import (
	"log"
	"runtime/debug"

	"github.com/joho/godotenv"
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
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("load .env file: %v", err)
	}
}
