package handlers

import "net/http"

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		get(w, r)
	} else if r.Method == http.MethodPost {
		post(w, r)
	} else if r.Method == http.MethodPut {
		put(w, r)
	} else if r.Method == http.MethodDelete {
		delete(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
