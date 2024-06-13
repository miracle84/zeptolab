package database

import (
	"context"
	"database/sql"
)

func RunMigrations(ctx context.Context, db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS file_data (
        id SERIAL PRIMARY KEY,
        type VARCHAR(15) NOT NULL,
        version VARCHAR(15) NOT NULL,
        hash VARCHAR(64) NOT NULL,
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    );
    `
	_, err := db.ExecContext(ctx, query)
	return err
}

func SaveData(ctx context.Context, db *sql.DB, fileType, version, hash, content string) error {
	query := "INSERT INTO file_data (type, version, hash, content) VALUES ($1, $2, $3, $4)"
	_, err := db.ExecContext(ctx, query, fileType, version, hash, content)
	return err
}
