package handler

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
	"plugin-file-content/model"
)

type MockLogger struct{}

func (m *MockLogger) Debug(format string, v ...interface{})                   {}
func (m *MockLogger) Info(format string, v ...interface{})                    {}
func (m *MockLogger) Warn(format string, v ...interface{})                    {}
func (m *MockLogger) Error(format string, v ...interface{})                   {}
func (m *MockLogger) WithFields(fields map[string]interface{}) runtime.Logger { return m }
func (m *MockLogger) WithField(key string, v interface{}) runtime.Logger      { return m }
func (m *MockLogger) Fields() map[string]interface{}                          { return map[string]interface{}{} }

type MockNakamaModule struct {
	//mock.Mock
	runtime.NakamaModule
}

var mockGetFilePath = func(basePath, fileType, version string) string {
	return "data/mock/path/to/file.json"
}

var mockReadFile = func(filePath string) ([]byte, error) {
	return []byte("testcontent"), nil
}

func TestRPCFunction(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	ctx := context.Background()
	logger := &MockLogger{}
	nk := &MockNakamaModule{}

	// Set up the test data
	payload := model.Payload{
		Type:    "core",
		Version: "1.0.0",
	}
	payloadBytes, err := json.Marshal(payload)
	assert.NoError(t, err)

	mock.ExpectExec("INSERT INTO file_data").
		WithArgs(payload.Type, payload.Version, "25edaa1f62bd4f2a7e4aa7088cf4c93449c1881af03434bfca027f1f82d69dba", "testcontent").
		WillReturnResult(sqlmock.NewResult(1, 1))

	ReadFile = mockReadFile
	GetFilePath = mockGetFilePath

	resp, err := FileContent(ctx, logger, mockDB, nk, string(payloadBytes))
	assert.NoError(t, err)

	var response model.Response
	err = json.Unmarshal([]byte(resp), &response)
	assert.NoError(t, err)
	assert.Equal(t, payload.Type, response.Type)
	assert.Equal(t, payload.Version, response.Version)
	assert.Equal(t, "25edaa1f62bd4f2a7e4aa7088cf4c93449c1881af03434bfca027f1f82d69dba", response.Hash)
	assert.Equal(t, "testcontent", response.Content)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
