package trivia

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"html"
)

// opentdbResponse models the API's top-level JSON response
type opentdbResponse struct {
	ResponseCode int              `json:"response_code"`
	Results      []opentdbQuestion `json:"results"`
}

// opentdbQuestion models each question from the API
type opentdbQuestion struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// FetchQuestions retrieves multiple-choice questions from Open Trivia DB
func FetchQuestions(amount int) ([]Question, error) {
	url := fmt.Sprintf("https://opentdb.com/api.php?amount=%d&type=multiple", amount)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch questions: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Parse JSON
	var apiResp opentdbResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Convert API results into our Question struct
	var questions []Question
for _, q := range apiResp.Results {
	question := Question{
		Text:             html.UnescapeString(q.Question),
		CorrectAnswer:    html.UnescapeString(q.CorrectAnswer),
		IncorrectAnswers: unescapeList(q.IncorrectAnswers),
		Category:         html.UnescapeString(q.Category),
		Difficulty:       html.UnescapeString(q.Difficulty),
	}
		question.Prepare()
		questions = append(questions, question)
	}

	return questions, nil
}

// unescapeList is a helper to decode HTML entities in answers
func unescapeList(list []string) []string {
	result := make([]string, len(list))
	for i, s := range list {
		result[i] = html.UnescapeString(s)
	}
	return result
}
