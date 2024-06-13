package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayload(t *testing.T) {
	payload := Payload{
		Type:    "core",
		Version: "1.0.0",
		Hash:    "testhash",
	}

	assert.Equal(t, "core", payload.Type)
	assert.Equal(t, "1.0.0", payload.Version)
	assert.Equal(t, "testhash", payload.Hash)
}

func TestResponse(t *testing.T) {
	response := Response{
		Type:    "core",
		Version: "1.0.0",
		Hash:    "testhash",
		Content: "testcontent",
	}

	assert.Equal(t, "core", response.Type)
	assert.Equal(t, "1.0.0", response.Version)
	assert.Equal(t, "testhash", response.Hash)
	assert.Equal(t, "testcontent", response.Content)
}
