package database

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunMigrations(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	ctx := context.Background()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS file_data").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = RunMigrations(ctx, mockDB)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestSaveData(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	ctx := context.Background()

	mock.ExpectExec("INSERT INTO file_data").
		WithArgs("core", "1.0.0", "testhash", "testcontent").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = SaveData(ctx, mockDB, "core", "1.0.0", "testhash", "testcontent")
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
