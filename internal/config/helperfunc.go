package config

import (
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(path, ".gatorconfig.json")
	return fullPath, nil
}