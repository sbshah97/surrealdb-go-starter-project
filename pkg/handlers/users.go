package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/gorilla/mux"
	"github.com/sbshah97/surrealdb-go-starter-project/pkg/helper"
	"github.com/sbshah97/surrealdb-go-starter-project/pkg/models"
)

func (h handler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse the JSON data from the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error("Error in unmarshalling POST request")
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	createdUserData, err := h.db.Create("users", user)
	if err != nil {
		// TODO: Remove panic and handle response usin http writer in handler
		slog.Error("Error in creating user data", "error", err)
		http.Error(w, fmt.Sprintf("Error in creating user data: %+v", err), http.StatusInternalServerError)
		return
	}

	created, err := helper.UnmarshalUsers(createdUserData)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in unmarshalling user data", "error", err)
		http.Error(w, fmt.Sprintf("Error in unmarshalling user data: %+v", err.Error()), http.StatusInternalServerError)
	}

	slog.Info("Successful response for POST Request")

	w.WriteHeader(http.StatusOK)
	// TODO: Create a successful response object
	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		slog.Error("Error in encoding response", "error", err)
		http.Error(w, fmt.Sprintf("Error in encoding response: %+v", err.Error()), http.StatusInternalServerError)
	}
}

func (h handler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	readUser, err := h.db.Fetch(vars["id"])
	if err != nil {
		// TODO: Remove panic and handle response usin http writer in handler
		slog.Error("Error in reading user data", "error", err)
		panic(err)
	}

	userObject, err := helper.UnmarshalSelectUserById(readUser)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in unmarshalling user data", "error", err, "readUser", readUser)
		panic(err)
	}
	slog.Info("After marshalling user data", "user", userObject)

	jsonObject, err := json.Marshal(userObject)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in marshalling JSON object", "error", err)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonObject); err != nil {
		slog.Error("Error writing response", "error", err)
	}
}
