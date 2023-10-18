package handlers

import (
	"net/http"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/database"
)

const (
	userTable = "users" // TODO: Refactor to a separate package
)

// webService is used to handle web requests via it's public methods
type handler struct {
	db *database.Database
}

func NewHandler(newDb *database.Database) *handler {
	return &handler{
		db: newDb,
	}
}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// TODO: Clean this up later
		// get(w, r)
	} else if r.Method == http.MethodPost {
		// TODO: Clean this up later
		// post(w, r)
	} else if r.Method == http.MethodPut {
		put(w, r)
	} else if r.Method == http.MethodDelete {
		delete(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
