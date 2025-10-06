package trivia

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestFullGameFlow(t *testing.T) {
	// --- Setup temporary config and highscore paths ---
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")
	highscorePath := filepath.Join(tmpDir, "highscore.json")

	// Write a temporary config file with a secret key
	cfg := ConfigData{SecretKey: "IntegrationKey"}
	cfgBytes, _ := json.Marshal(cfg)
	if err := os.WriteFile(configPath, cfgBytes, 0644); err != nil {
		t.Fatalf("Failed to write temp config: %v", err)
	}

	// Load the secret key
	key, err := loadSecretKeyFromPath(configPath)
	if err != nil {
		t.Fatalf("Failed to load secret key from temp config: %v", err)
	}

	// --- Create and prepare a sample question ---
	q := &Question{
		Text:            "Capital of France?",
		CorrectAnswer:   "Paris",
		IncorrectAnswers: []string{"London", "Rome", "Berlin"},
		Category:        "Geography",
		Difficulty:      "Easy",
	}
	q.Prepare()

	if len(q.ShuffledChoices) != 4 {
		t.Fatalf("Expected 4 choices, got %d", len(q.ShuffledChoices))
	}

	// --- Simulate answering correctly ---
	correctLabel := []string{"A", "B", "C", "D"}[q.CorrectIndex]
	if !q.CheckAnswer(correctLabel) {
		t.Errorf("Correct answer check failed")
	}

	// --- Save new highscore ---
	name := "Tester"
	score := 1 // since answered one question correctly
	if err := SaveHighScoreWithKey(highscorePath, name, score, key); err != nil {
		t.Fatalf("SaveHighScoreWithKey failed: %v", err)
	}

	// --- Load highscore and verify ---
	loadedName, loadedScore, err := LoadHighScoreWithKey(highscorePath, key)
	if err != nil {
		t.Fatalf("LoadHighScoreWithKey failed: %v", err)
	}
	if loadedName != name || loadedScore != score {
		t.Errorf("Expected loaded highscore (%s, %d), got (%s, %d)", name, score, loadedName, loadedScore)
	}
}
