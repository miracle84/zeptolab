package file

import (
	"os"
	"path/filepath"
)

func GetFilePath(BasePath, fileType, version string) string {
	return filepath.Join(BasePath, fileType, version+".json")
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
