package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
