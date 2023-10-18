package database

import "log/slog"

func (d Database) Fetch(what string) (interface{}, error) {
	// Insert user
	data, err := d.db.Select(what)
	if err != nil {
		slog.Error("Unable to insert data in table: %v", err)
		return nil, err
	}

	slog.Debug("Successfully fetched data in table")

	return data, nil
}
