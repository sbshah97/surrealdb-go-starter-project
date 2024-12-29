package database

import "log/slog"

func (d Database) Update(userID string, data interface{}) (interface{}, error) {
	slog.Debug("ID passed to update", "UserID", userID)

	// Update user by ID passed
	vars := make(map[string]interface{})
	vars["userID"] = userID
	updatedData, err := d.db.Query("UPDATE type::thing(\"users\", $userID) CONTENT $data", map[string]interface{}{
		"userID": userID,
		"data":   data,
	})
	if err != nil {
		slog.Error("Unable to update data in table", "error", err)
		return nil, err
	}

	slog.Info("Successfully updated data in table", "data", updatedData)

	return updatedData, nil
}
