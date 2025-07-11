package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

type DBConfig struct {
	Provider string
	Driver   string
	DBPath   string
}

func Open(dbConfig *DBConfig) (*sql.DB, error) {
	// Ensure the parent directory exists
	if err := os.MkdirAll(filepath.Dir(dbConfig.DBPath), 0755); err != nil {
		return nil, fmt.Errorf("db: mkdir %w", err)
	}

	db, err := sql.Open(dbConfig.Driver, dbConfig.DBPath)
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db: ping %w", err)
	}

	fmt.Printf(`----- Connected to Database -----
Provider: %s
DB: %s
---------------------------------
`,
		dbConfig.Provider,
		dbConfig.DBPath,
	)

	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("sqlite")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
