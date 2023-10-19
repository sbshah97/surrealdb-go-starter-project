package database

import (
	"log/slog"
)

func (d Database) Create(thing string, data interface{}) (interface{}, error) {
	// Insert user
	slog.Debug("Thing", thing, "data", data)

	data, err := d.db.Create(thing, data)
	if err != nil {
		slog.Error("Unable to insert data in table: %v", err)
		return nil, err
	}

	slog.Debug("Successfully inserted data in table")

	return data, nil
}
