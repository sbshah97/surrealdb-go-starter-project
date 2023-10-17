package helper

import "github.com/sbshah97/surrealdb-go-starter-project/pkg/models"

func PopulateUserData() models.User {
	// Create user struct
	user := models.User{
		Name:    "Salman",
		Surname: "Shah",
	}

	return user
}
