package main

func main() {
    res, _ := http.Get("https://opentdb.com/api.php?amount=50&category=18&difficulty=medium")

    triviaRes := TriviaResponse{}

    bytes, _ := io.ReadAll(res.Body)

    json.Unmarshal(bytes, &triviaRes)

    for _, result := range triviaRes.REsult {
        printLn(result.Category + ": " + result.Question)
    }
}
