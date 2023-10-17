package handlers

import (
	"net/http"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/helper"
	"golang.org/x/exp/slog"
)

func (h handler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createdUserData, err := h.db.Create("users", helper.PopulateUserData())
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

		w.WriteHeader(http.StatusOK)
		slog.Info("Successful response")
		// TODO: Create a successful response object
		// w.Write("Successful response")
	}
	w.WriteHeader(http.StatusBadRequest)
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		put(w, r)
	}
}

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		get(w, r)
	}
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		delete(w, r)
	}
}
