package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name        string
		packSizes   []int
		expected    []int
		expectError bool
		invalidJSON bool
		missingFile bool
	}{
		{
			name:        "Valid config returns sorted descending",
			packSizes:   []int{500, 250, 1000, 2000, 5000},
			expected:    []int{5000, 2000, 1000, 500, 250},
			expectError: false,
		},
		{
			name:        "Empty pack sizes returns empty slice",
			packSizes:   []int{},
			expected:    []int{},
			expectError: false,
		},
		{
			name:        "Invalid JSON returns error",
			invalidJSON: true,
			expectError: true,
		},
		{
			name:        "Missing file returns error",
			missingFile: true,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filename string

			if tt.missingFile {
				// Point to a non-existent file
				filename = filepath.Join(t.TempDir(), "nonexistent.json")
			} else {
				// Create a temp file
				tmpDir := t.TempDir()
				filename = filepath.Join(tmpDir, "config.json")

				var content []byte
				if tt.invalidJSON {
					content = []byte(`{ invalid json }`)
				} else {
					jsonData, err := json.Marshal(Config{PackSizes: tt.packSizes})
					assert.NoError(t, err, "JSON marshal failed in test setup")
					content = jsonData
				}

				err := os.WriteFile(filename, content, 0644)
				assert.NoError(t, err, "failed to write test config file")
			}

			result, err := LoadConfig(filename)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result, "pack sizes mismatch")
			}
		})
	}
}
