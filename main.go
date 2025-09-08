package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getQuestions() ([]Question, error) {
    res, err := http.Get("https://opentdb.com/api.php?amount=10")
    if err != nil {
        log.Printf("Error fetching questions: %v\n", err)
        return nil, err
    }

    defer res.Body.Close()
    var questionsResponse QuestionsResponse
    err = json.NewDecoder(res.Body).Decode(&questionsResponse)
    if err != nil {
        log.Printf("Error decoding response: %v\n", err)
        return nil, err
    }
    return questionsResponse.Results, nil
}

func main() {
    questions, err := getQuestions()
    if err != nil {
        log.Fatalf("Failed to get questions: %v", err)
    }
    for _, q := range questions {
        log.Printf("Question: %s\n", q.Question)
    }
}
