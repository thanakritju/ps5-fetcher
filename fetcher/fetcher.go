package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"ps5-fetcher/line"
	"regexp"
	"strings"
)

type FetcherService struct {
	targetUrls   []string
	keywordRegex *regexp.Regexp
	lineService  *line.LineService
}

func NewFetcherService(lineService *line.LineService) *FetcherService {
	targetUrls := strings.Split(os.Getenv("TARGET_URLS"), ",")
	keywordRegex := regexp.MustCompile(os.Getenv("KEYWORD_REGEX"))
	return &FetcherService{targetUrls: targetUrls, keywordRegex: keywordRegex, lineService: lineService}
}

func (s FetcherService) Run() {
	for _, url := range s.targetUrls {
		fmt.Printf("making http request to: %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error parsing http response: %s\n", err)
			os.Exit(1)
		}

		if s.keywordRegex.Match(body) {
			message := fmt.Sprintf("Found matching at: %s", url)
			s.lineService.SendMessage(message)
		}
	}
}