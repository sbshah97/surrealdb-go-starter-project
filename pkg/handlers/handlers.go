package handlers

import (
	"net/http"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/database"
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

// TODO: Implement in the future as a generic method handler
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// TODO: Clean this up later
		// get(w, r)
	} else if r.Method == http.MethodPost {
		// TODO: Clean this up later
		// post(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
