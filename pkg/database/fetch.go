package database

import "log/slog"

func (d Database) Fetch(what string) (interface{}, error) {
	slog.Debug("ID passed to select", what)

	// Select user by ID passed
	data, err := d.db.Select(what)
	if err != nil {
		slog.Error("Unable to insert data in table: %v", err)
		return nil, err
	}

	slog.Info("Successfully fetched data in table", data)

	return data, nil
}
