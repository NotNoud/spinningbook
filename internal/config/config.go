package config

import (
	"os"
	"path/filepath"
)

func Dir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "spinningbook"), nil
}

func MacrosFile() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "macros.toml"), nil
}
