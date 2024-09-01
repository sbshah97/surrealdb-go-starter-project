package main

import (
	"net/http"
	"strconv"
	"time"

	"log/slog"

	"github.com/gorilla/mux"

	"github.com/sbshah97/surrealdb-go-starter-project/pkg/database"
	"github.com/sbshah97/surrealdb-go-starter-project/pkg/handlers"
)

const (
	wsString  = "ws://db:8000/rpc"
	tableName = "user"
	namespace = "test"
	dbName    = "test"
	port      = 8080
	username  = "root"
	password  = "root"
)

func main() {
	db, err := database.NewDatabase(wsString, username, password, namespace, dbName)
	if err != nil {
		slog.Error("failed to create shortener repository", "error", err)
		return
	}

	slog.Info("Connected to database")
	// defer database.DB.Close()
	// Close connections to the database at program shutdown
	defer func() {
		slog.Info("Closing database")
		db.Close()
	}()

	handler := handlers.NewHandler(db)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.FetchUsers).Methods("GET", "PUT")
	r.HandleFunc("/users", handler.CreateUsers).Methods("POST")
	http.Handle("/", r)

	slog.Info("Listening on http://localhost:" + strconv.Itoa(port))
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		slog.Error("failed to listen", "error", err)
	}

}
