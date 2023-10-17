package main

import (
	"github.com/sbshah97/surrealdb-go-starter-project/pkg/database"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/helper"
	"github.com/surrealdb/surrealdb.go"
)

var db *surrealdb.DB

const (
	wsString  = "ws://localhost:8000/rpc"
	tableName = "user"
	namespace = "test"
	dbName    = "test"
)

func init() {
	var err error
	// Connect to SurrealDB
	db, err = surrealdb.New(wsString)
	if err != nil {
		panic(err)
	}

	// TODO: Sign in if env is prod
	// Sign in
	// if _, err = db.Signin(map[string]string{
	// 	"user": "root",
	// 	"pass": "root",
	// }); err != nil {
	// 	panic(err)
	// }

	// Select namespace and database
	if _, err := db.Use(namespace, dbName); err != nil {
		panic(err)
	}
}

func main() {
	createdUserData, err := database.Create(db, tableName, helper.PopulateUserData())
	if err != nil {
		// TODO: Remove panic and add slog
		panic(err)
	}

	_, err = helper.UnmarshalUser(createdUserData)
	if err != nil {
		// TODO: Remove panic and add slog
		panic(err)
	}

	// // Get user by ID
	// data, err = db.Select(createdUser[0].ID)
	// if err != nil {
	// 	panic(err)
	// }

	// // Unmarshal data
	// selectedUser := new(User)
	// err = surrealdb.Unmarshal(data, &selectedUser)
	// if err != nil {
	// 	panic(err)
	// }

	// // Change part/parts of user
	// changes := map[string]string{"name": "Jane"}
	// if _, err = db.Change(selectedUser.ID, changes); err != nil {
	// 	panic(err)
	// }

	// // Update user
	// if _, err = db.Update(selectedUser.ID, changes); err != nil {
	// 	panic(err)
	// }

	// // Raw Query user
	// if _, err = db.Query("SELECT * FROM $record", map[string]interface{}{
	// 	"record": createdUser[0].ID,
	// }); err != nil {
	// 	panic(err)
	// }

	// // Delete user by ID
	// if _, err = db.Delete(selectedUser.ID); err != nil {
	// 	panic(err)
	// }
}
