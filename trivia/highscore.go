package trivia

import (
	"encoding/json"
	"fmt"
	"os"
)

// HighScoreData represents the stored high score with player's name
type HighScoreData struct {
	Name      string `json:"name"`
	HighScore int    `json:"highscore"`
}

// LoadHighScore reads the high score from a file
func LoadHighScore(filename string) (string, int) {
	plain, err := os.ReadFile(filename)
	if err != nil {
		// file missing -> no highscore yet
		return "", 0
	}	
	
// parse JSON
	var data HighScoreData
	if err := json.Unmarshal(plain, &data); err != nil {
		return "", 0
	}

	return data.Name, data.HighScore
}

// SaveHighScore writes the high score to a file
func SaveHighScore(filename string, name string, score int) error {
	data := HighScoreData{
		Name:      name,
		HighScore: score,
	}
	plain, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal highscore json: %w", err)
	}
	// write to file
	if err := os.WriteFile(filename, []byte(plain), 0644); err != nil {
		return fmt.Errorf("failed to write highscore file: %w", err)
	}

	return nil
}
