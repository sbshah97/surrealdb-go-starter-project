package database

import "log/slog"

func (d Database) Fetch(userID string) (interface{}, error) {
	slog.Debug("ID passed to select", userID)

	// Select user by ID passed
	vars := make(map[string]interface{})
	vars["userID"] = userID
	data, err := d.db.Query("SELECT * FROM type::thing(\"users\", $userID)", vars)
	if err != nil {
		slog.Error("Unable to read data from table", "error", err)
		return nil, err
	}

	slog.Info("Successfully fetched data in table", "data", data)

	return data, nil
}
