package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	expectedPath := filepath.Join("data", "core", "1.0.0.json")
	actualPath := GetFilePath("data", "core", "1.0.0")
	assert.Equal(t, expectedPath, actualPath)
}

func TestReadFile(t *testing.T) {
	// Create a temporary file
	filePath := filepath.Join(os.TempDir(), "testfile")
	expectedContent := []byte("test content")
	err := os.WriteFile(filePath, expectedContent, 0644)
	assert.NoError(t, err)
	defer os.Remove(filePath)

	actualContent, err := ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, actualContent)
}
