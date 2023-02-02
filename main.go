package main

import (
	"log"
	"ps5-fetcher/fetcher"
	"ps5-fetcher/line"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lineService := line.NewLineService()
	fetcher := fetcher.NewFetcherService(lineService)
	fetcher.Run()
}
