package main

import (
	"context"
	"database/sql"
	"log"
	"plugin-file-content/database"
	"plugin-file-content/handler"

	"github.com/heroiclabs/nakama-common/runtime"
)

const BasePath = "data"

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterRpc("fileContent", handler.FileContent); err != nil {
		return err
	}

	logger.Info("Running migrations...")
	if err := database.RunMigrations(ctx, db); err != nil {
		return err
	}
	logger.Info("Migrations completed successfully")

	return nil
}

func main() {
	log.Println("Nakama module loaded.")
}
