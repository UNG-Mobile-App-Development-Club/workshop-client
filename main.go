package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("https://opentdb.com/api.php?amount=10&difficulty=hard")

	triviaRes := TriviaResponse{}

	bytes, _ := io.ReadAll(res.Body)

	json.Unmarshal(bytes, &triviaRes)

	for _, result := range triviaRes.Results {
		println(result.Category + ": " + result.Question)
	}
}
