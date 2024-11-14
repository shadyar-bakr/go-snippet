package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/shadyar-bakr/go-snippet/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := gorm.Open(sqlite.Open("snippets.db"), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("failed to get database instance", "error", err)
		os.Exit(1)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	defer sqlDB.Close()

	app := &application{
		logger:   logger,
		snippets: models.NewSnippetModel(db),
	}

	logger.Info("Server Started", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
