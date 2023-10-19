package database

import (
	"fmt"

	"log/slog"

	"github.com/surrealdb/surrealdb.go"
)

type Database struct {
	db *surrealdb.DB
}

type DatabaseParams struct {
	Address   string
	User      string
	Password  string
	Namespace string
	Database  string
}

// TODO: Use database params instead of passing parameters individually
func NewDatabase(address, user, password, namespace, database string) (*Database, error) {
	//surreal.New is
	db, err := surrealdb.New(address)
	if err != nil {
		slog.Error("Unable to initialise SurrealDB", err)
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	// TODO: Enable for non production instance
	_, err = db.Signin(map[string]interface{}{
		"user": user,
		"pass": password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sign in: %w", err)
	}

	_, err = db.Use(namespace, database)
	if err != nil {
		slog.Error("Unable to use DB", err)
		return nil, err
	}

	return &Database{db}, nil
}

func (d Database) Close() {
	d.db.Close()
}
