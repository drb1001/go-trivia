package trivia

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// HighScoreData represents the stored high score with player's name
type HighScoreData struct {
	Name      string `json:"name"`
	HighScore int    `json:"highscore"`
}

// ConfigData represents local config stored in config.json
type ConfigData struct {
	SecretKey string `json:"secret_key"`
}

const (
	configFileName   = "config.json"
)

// getConfigPath returns the path to config.json
func getConfigPath() (string, error) {
	return filepath.Join(configFileName), nil
}

// GetConfigPathForUser returns the path as a string and a nil error when possible.
// Exposed so main can show helpful instructions in case of missing config.
func GetConfigPathForUser() (string, error) {
	return getConfigPath()
}

// loadSecretKey loads the obfuscation key from config.json
// Returns an error if the config file does not exist or is invalid.
func loadSecretKey() ([]byte, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file not found at %s", configPath)
		}
		return nil, fmt.Errorf("failed to read config file at %s: %w", configPath, err)
	}

	var cfg ConfigData
	if err := json.Unmarshal(content, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config file at %s: %w", configPath, err)
	}
	if cfg.SecretKey == "" {
		return nil, fmt.Errorf("config file at %s contains empty secret_key", configPath)
	}

	return []byte(cfg.SecretKey), nil
}

// LoadHighScore reads the (obfuscated) high score file and returns name and score.
// Returns an error if the config/key is missing or the highscore file is unreadable.
func LoadHighScore(filename string) (string, int, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		// If file doesn't exist, that's okay — return zero score but still require key presence.
		if os.IsNotExist(err) {
			// still require a valid key to proceed
			if _, keyErr := loadSecretKey(); keyErr != nil {
				return "", 0, keyErr
			}
			return "", 0, nil
		}
		return "", 0, fmt.Errorf("failed to read highscore file: %w", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		return "", 0, fmt.Errorf("failed to base64-decode highscore file: %w", err)
	}

	key, err := loadSecretKey()
	if err != nil {
		return "", 0, err
	}

	plain := xorBytes(decoded, key)

	var data HighScoreData
	if err := json.Unmarshal(plain, &data); err != nil {
		return "", 0, fmt.Errorf("failed to parse highscore data: %w", err)
	}

	return data.Name, data.HighScore, nil
}

// SaveHighScore writes the obfuscated highscore file with name and score.
func SaveHighScore(filename string, name string, score int) error {

	key, err := loadSecretKey()
	if err != nil {
		return err
	}

	data := HighScoreData{
		Name:      name,
		HighScore: score,
	}

	plain, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal highscore json: %w", err)
	}

	obf := xorBytes(plain, key)
	enc := base64.StdEncoding.EncodeToString(obf)

	if err := os.WriteFile(filename, []byte(enc), 0644); err != nil {
		return fmt.Errorf("failed to write highscore file: %w", err)
	}

	return nil
}

// xorBytes XORs input bytes with key repeating as necessary.
func xorBytes(data, key []byte) []byte {
	if len(key) == 0 {
		return data
	}
	out := make([]byte, len(data))
	for i := range data {
		out[i] = data[i] ^ key[i%len(key)]
	}
	return out
}
