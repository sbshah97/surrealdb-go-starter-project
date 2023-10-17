package handlers

import "net/http"

func post(w http.ResponseWriter, r *http.Request) {
	// createdUserData, err := database.Create(db, tableName, helper.PopulateUserData())
	// if err != nil {
	// 	// TODO: Remove panic and add slog
	// 	panic(err)
	// }

	// _, err = helper.UnmarshalUser(createdUserData)
	// if err != nil {
	// 	// TODO: Remove panic and add slog
	// 	panic(err)
	// }
	w.WriteHeader(http.StatusBadRequest)
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
