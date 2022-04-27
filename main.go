package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	println("listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handleTransactions(r *mux.Router) {
	s := r.PathPrefix("/api/transaction").Subrouter()
	s.Use(AuthMiddleware)

	s.HandleFunc("/all", api.GetTransactions)
	s.HandleFunc("/add", api.CreateTransaction).Methods("POST")
	s.HandleFunc("/update/{id}", api.UpdateTransaction).Methods("POST")
	s.HandleFunc("/delete/{id}", api.DeleteTransaction).Methods("DELETE")
}

// Handle Basic Auth for every route this middleware is added to.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if BasicAuth(w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

func BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	username := os.Getenv("HTTP_USERNAME")
	password := os.Getenv("HTTP_PASSWORD")
	user, pass, ok := r.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))
		return false
	}

	return true
}
