package main

type TriviaResponse struct {
	ResponseCode int      `json:"resonse_code"`
	Results      []Result `json:"results"`
}

type Result struct {
	Question string `json:"question"`
	Category string `json:"category"`
}
