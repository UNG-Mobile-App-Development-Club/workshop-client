package main

type TriviaResponse struct {
	ResponseCode int `json:"response_code"`
	Results      []Result `json:"results"`
}

type Result struct {
	Question string `json:"question"`
	Category string `json:"category"`

}
