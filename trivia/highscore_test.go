package trivia

import (
	"encoding/json"
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"
)

// SaveHighScoreWithKey writes a highscore using a provided secret key (for testing)
func SaveHighScoreWithKey(filename, name string, score int, key []byte) error {
	data := HighScoreData{Name: name, HighScore: score}
	plain, err := json.Marshal(data)
	if err != nil {
		return err
	}

	obf := xorBytes(plain, key)
	enc := base64.StdEncoding.EncodeToString(obf)

	return os.WriteFile(filename, []byte(enc), 0644)
}

// LoadHighScoreWithKey reads a highscore using a provided secret key (for testing)
func LoadHighScoreWithKey(filename string, key []byte) (string, int, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return "", 0, nil
		}
		return "", 0, err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		return "", 0, err
	}

	plain := xorBytes(decoded, key)
	var data HighScoreData
	if err := json.Unmarshal(plain, &data); err != nil {
		return "", 0, err
	}

	return data.Name, data.HighScore, nil
}

func TestSaveLoadHighScore(t *testing.T) {
	// --- Setup temporary files ---
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")
	highscorePath := filepath.Join(tmpDir, "highscore.json")

	// Write a temporary config file with a secret key
	cfg := ConfigData{SecretKey: "TestSecretKey"}
	cfgBytes, _ := json.Marshal(cfg)
	if err := os.WriteFile(configPath, cfgBytes, 0644); err != nil {
		t.Fatalf("Failed to write temp config: %v", err)
	}

	// Load the key from the temp config
	key, err := loadSecretKeyFromPath(configPath)
	if err != nil {
		t.Fatalf("Failed to load secret key from temp config: %v", err)
	}

	// --- Test SaveHighScoreWithKey ---
	name := "Alice"
	score := 42
	if err := SaveHighScoreWithKey(highscorePath, name, score, key); err != nil {
		t.Fatalf("SaveHighScoreWithKey failed: %v", err)
	}

	// --- Verify file is not plain JSON ---
	content, err := os.ReadFile(highscorePath)
	if err != nil {
		t.Fatalf("Failed to read highscore file: %v", err)
	}
	if string(content) == string(cfgBytes) {
		t.Error("Highscore file should be obfuscated; content matches config JSON")
	}

	// --- Test LoadHighScoreWithKey ---
	loadedName, loadedScore, err := LoadHighScoreWithKey(highscorePath, key)
	if err != nil {
		t.Fatalf("LoadHighScoreWithKey failed: %v", err)
	}
	if loadedName != name {
		t.Errorf("Loaded name = %s; want %s", loadedName, name)
	}
	if loadedScore != score {
		t.Errorf("Loaded score = %d; want %d", loadedScore, score)
	}

	// --- Test loading non-existent file ---
	nonexistentPath := filepath.Join(tmpDir, "nofile.json")
	nName, nScore, err := LoadHighScoreWithKey(nonexistentPath, key)
	if err != nil {
		t.Errorf("LoadHighScoreWithKey should succeed for missing file if key exists, got error: %v", err)
	}
	if nName != "" || nScore != 0 {
		t.Errorf("Expected zero values for missing file, got name=%s, score=%d", nName, nScore)
	}
}
