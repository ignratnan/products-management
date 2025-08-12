package forJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJson(folderpath string, filename string, data interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal the data: %w", err)
	}

	err = os.WriteFile(fullpath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write the file: %w", err)
	}

	return nil
}

func ReadJson(folderpath string, filename string, target interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := os.ReadFile(fullpath)
	if err != nil {
		return fmt.Errorf("failed to read the file: %w", err)
	}

	err = json.Unmarshal(jsonData, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the data: %w", err)
	}

	return nil
}
