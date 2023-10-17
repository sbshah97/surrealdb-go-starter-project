package helper

import (
	"fmt"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/models"
	"github.com/surrealdb/surrealdb.go"
)

func UnmarshalUser(data interface{}) ([]models.User, error) {
	// Unmarshal data
	createdUser := make([]models.User, 1)
	err := surrealdb.Unmarshal(data, &createdUser)
	if err != nil {
		return createdUser, err
	}

	fmt.Println("Created User in DB", createdUser)

	return createdUser, nil
}
