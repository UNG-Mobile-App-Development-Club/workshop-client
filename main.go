package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://opentdb.com/api.php?amount=10&difficulty=hard")
	if err != nil {
		log.Fatalf("Failed to fetch trivia questions: %v", err)
	}

	triviaRes := TriviaResponse{}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	json.Unmarshal(bytes, &triviaRes)

	for _, result := range triviaRes.Results {
		println(result.Category + ": " + result.Question)
	}
}
