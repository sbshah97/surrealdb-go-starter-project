package handlers

import (
	"encoding/json"
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
		slog.Error("Error in creating user data: %v", err)
		http.Error(w, "Error in creating user data: Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = helper.UnmarshalUsers(createdUserData)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in unmarshalling user data: %v", err)
		http.Error(w, "Error in unmarshalling user data: Error"+err.Error(), http.StatusInternalServerError)
	}

	slog.Info("Successful response for POST Request")

	w.WriteHeader(http.StatusOK)
	// TODO: Create a successful response object
	// w.Write("Successful response")
	return

}

func (h handler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	createdUserData, err := h.db.Fetch("users:" + vars["id"])
	if err != nil {
		// TODO: Remove panic and handle response usin http writer in handler
		slog.Error("Error in creating user data: %v", err)
		panic(err)
	}

	userObject, err := helper.UnmarshalSelectUserById(createdUserData)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in creating user data: %v", err)
		panic(err)
	}

	jsonObject, err := json.Marshal(userObject)
	if err != nil {
		// TODO: Remove panic and add slog
		slog.Error("Error in marshalling JSON object: %v", err)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonObject)
	return
}
