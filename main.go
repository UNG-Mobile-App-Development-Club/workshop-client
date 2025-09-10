package main

import "net/http"
import "io"
import "encoding/json"

func main() {
    res, _ := http.Get("https://opentdb.com/api.php?amount=10&category=18&difficulty=medium")

        triviaRes := TriviaResponse{}

        bytes, _ := io.ReadAll(res.Body)

        json.Unmarshal(bytes, &triviaRes)

        for _, result := range triviaRes.Results {
            println(result.Category + ": " + result.Question)
        }
}