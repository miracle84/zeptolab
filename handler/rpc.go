package handler

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"plugin-file-content/model"

	"github.com/heroiclabs/nakama-common/runtime"
	"plugin-file-content/database"
	"plugin-file-content/file"
)

var GetFilePath = file.GetFilePath
var ReadFile = file.ReadFile

// Function was selected just to calculate hash,
// possibly need to select it more carefully based on performance
// and other requirements
var hashFunc = func(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}

var basePath = "data"

func FileContent(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	var p model.Payload
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		return "", err
	}

	// Set default values if necessary
	if p.Type == "" {
		p.Type = model.DefaultType
	}
	if p.Version == "" {
		p.Version = model.DefaultVersion
	}

	// Need to add Validation to prevent access outside of dir or inappropriate values for version or type
	// Read file content
	filePath := GetFilePath(basePath, p.Type, p.Version)
	fileContent, err := ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("file %s does not exist", filePath)
	}

	fileHash := hashFunc(fileContent)

	// Prepare response
	response := model.Response{
		Type:    p.Type,
		Version: p.Version,
		Hash:    fileHash,
	}

	// Check hash
	if p.Hash != "" && p.Hash != fileHash {
		response.Content = ""
	} else {
		response.Content = string(fileContent)
	}

	// Save data to database
	if err := database.SaveData(ctx, db, p.Type, p.Version, fileHash, response.Content); err != nil {
		return "", err
	}

	// Convert response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(responseJSON), nil
}
