package helper

import (
	"fmt"
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
	arr, ok := data.([]interface{})
	if !ok {
		return models.User{}, fmt.Errorf("Expected the data to be an array but instead was %T", data)
	}
	if len(arr) != 1 {
		return models.User{}, fmt.Errorf("Expected the data to be an array of length 1 but instead was %d", len(arr))
	}
	resultMap, ok := arr[0].(map[string]interface{})
	if !ok {
		return models.User{}, fmt.Errorf("Expected the data to be an object but instead was %T", arr[0])
	}
	userMapArrayIface, ok := resultMap["result"]
	if !ok {
		return models.User{}, fmt.Errorf("Expected the data to be an object with a'result' key but instead was %T", resultMap)
	}
	userMapArray, ok := userMapArrayIface.([]interface{})
	if !ok {
		return models.User{}, fmt.Errorf("Expected the data to be an array of objects but instead was %T", userMapArrayIface)
	}
	if len(userMapArray) != 1 {
		return models.User{}, fmt.Errorf("Expected the data to be an array of length 1 but instead was %d", len(userMapArray))
	}
	userMap, ok := userMapArray[0].(map[string]interface{})
	if !ok {
		return models.User{}, fmt.Errorf("Expected the data to be an object but instead was %T", userMapArrayIface)
	}
	user := models.User{}
	user.ID = userMap["id"].(string)
	user.Name = userMap["name"].(string)
	user.Surname = userMap["surname"].(string)

	slog.Info("Created object is", "user", user)
	return user, nil
}
