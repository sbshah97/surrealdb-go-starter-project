package handlers

import "net/http"

func put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
