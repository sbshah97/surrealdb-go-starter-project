package helper

import (
	"log/slog"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/models"
	"github.com/surrealdb/surrealdb.go"
)

func UnmarshalUsers(data interface{}) ([]models.User, error) {
	// Unmarshal data
	createdUser := make([]models.User, 1)
	err := surrealdb.Unmarshal(data, &createdUser)
	if err != nil {
		return createdUser, err
	}

	slog.Info("Created object is", createdUser)
	return createdUser, nil
}

func UnmarshalSelectUserById(data interface{}) (models.User, error) {
	// Unmarshal data
	selectedUser := new(models.User)
	err := surrealdb.Unmarshal(data, &selectedUser)
	if err != nil {
		return *selectedUser, err
	}

	slog.Info("Selected object is", *selectedUser)
	return *selectedUser, nil
}
