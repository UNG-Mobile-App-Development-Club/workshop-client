package main

type QuestionsResponse struct {
    ResponseCode int `json:"response_code"`
    Results      []Question `json:"results"`
}

type Question struct {
    Type                string   `json:"type"`
    Difficulty          string   `json:"difficulty"`
    Category            string   `json:"category"`
    Question            string   `json:"question"`
    CorrectAnswer       string   `json:"correct_answer"`
    IncorrectAnswers    []string `json:"incorrect_answers"`
}

