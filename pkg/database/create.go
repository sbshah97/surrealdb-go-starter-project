package database

import "github.com/surrealdb/surrealdb.go"

func Create(db *surrealdb.DB, thing string, data interface{}) (interface{}, error) {
	// Insert user
	data, err := db.Create(thing, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
