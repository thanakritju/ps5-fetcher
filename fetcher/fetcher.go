package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"ps5-fetcher/line"
	"regexp"
)

type FetcherService struct {
	targets     Targets
	lineService *line.LineService
}

type Targets struct {
	Targets []struct {
		URL     string `json:"url"`
		Pattern string `json:"pattern"`
	} `json:"targets"`
}

func NewFetcherService(lineService *line.LineService) *FetcherService {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var targets Targets
	json.Unmarshal(byteValue, &targets)
	log.Printf("Loaded config: %s\n", byteValue)
	return &FetcherService{targets: targets, lineService: lineService}
}

func (s FetcherService) Run() {
	for _, t := range s.targets.Targets {
		body := get(t.URL)

		match, _ := regexp.Match(t.Pattern, body)
		if match {
			message := fmt.Sprintf("Found matching at: %s", t.URL)
			s.lineService.SendMessage(message)
		}
	}
}

func get(url string) []byte {
	log.Printf("Making http request to: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making http request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error parsing http response: %s\n", err)
		os.Exit(1)
	}
	return body
}
