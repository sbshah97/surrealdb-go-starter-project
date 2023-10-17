package handlers

import (
	"log/slog"
	"net/http"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/helper"
)

func (h handler) post(w http.ResponseWriter, r *http.Request, tableName string) {
	createdUserData, err := h.db.Create(tableName, helper.PopulateUserData())
	if err != nil {
		// TODO: Remove panic and handle response usin http writer in handler
		slog.Error("Error in creating user data: %v", err)
		panic(err)
	}

	_, err = helper.UnmarshalUser(createdUserData)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in creating user data: %v", err)
		panic(err)
	}
}
