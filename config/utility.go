package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ParseEnv(env string, fallback string, fallbackEnabled bool) string {
	variable := os.Getenv(env)

	if fallbackEnabled && variable == "" {
		variable = fallback
	}
	if variable == "" {
		log.Fatalf("'%s' environment variable is required.", env)
	}

	return variable
}

func PurgeEmptyFolders(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %w", path, err)
		}

		if !info.IsDir() {
			return nil
		}

		files, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("error reading directory %s: %w", path, err)
		}

		if len(files) == 0 {
			err = os.Remove(path)
			if err != nil {
				return fmt.Errorf("error removing directory %s: %w", path, err)
			}
			fmt.Printf("Removed empty folder: %s\n", path)
		}

		return nil
	})
}
