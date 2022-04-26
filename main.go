package main

import (
	"log"
	"net/http"
	db "webapp/config/database"
	"webapp/src/api"

	"github.com/gorilla/mux"
)

func main() {
	if err := db.Open(); err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	handleTransactions(r)

	port := "8080"
	println("listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handleTransactions(r *mux.Router) {
	s := r.PathPrefix("/api/transaction").Subrouter()

	s.HandleFunc("/all", api.GetTransactions)
	s.HandleFunc("/add", api.CreateTransaction).Methods("POST")
	s.HandleFunc("/update/{id}", api.UpdateTransaction).Methods("POST")
	s.HandleFunc("/delete/{id}", api.DeleteTransaction).Methods("DELETE")
}
