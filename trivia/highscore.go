package trivia

import (
	"encoding/json"
	"fmt"
	"os"
)

// HighScoreData represents the JSON structure for high score file
type HighScoreData struct {
	HighScore int `json:"highscore"`
}

// LoadHighScore reads the high score from a file
func LoadHighScore(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		// If file doesn’t exist, return 0
		return 0
	}
	defer file.Close()

	var data HighScoreData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return 0
	}
	return data.HighScore
}

// SaveHighScore writes the high score to a file
func SaveHighScore(filename string, score int) {
	data := HighScoreData{HighScore: score}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error saving high score: %v\n", err)
		return
	}
	defer file.Close()

	_ = json.NewEncoder(file).Encode(data)
}
