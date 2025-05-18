package config

import (
	"encoding/json"
	"os"
)

// New loads a configuration from a JSON file.
// This is a minimal stub used for testing without the real library.
func New[T any](path string) (*T, error) {
	var cfg T
	if path == "" {
		return &cfg, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
