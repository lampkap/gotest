package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	db "webapp/config/database"

	"github.com/gorilla/mux"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	// Find all transactions in the database.
	transactions := []db.Transaction{}
	db.DB.Find(&transactions)
	// Return the transactions as json data.
	json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction db.Transaction

	// Handle incoming JSON data.
	if err := handleTransactionData(w, r, &transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the transaction into the database.
	result := db.DB.Create(&transaction)

	if result.Error != nil {
		panic(result.Error)
	}

	// Return newly created transaction.
	json.NewEncoder(w).Encode(transaction)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	// Get variable from route.
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if id is an integer.
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, "Id is invalid", http.StatusBadRequest)
		return
	}

	var transaction db.Transaction

	// Check if getting the transaction by id throws an error.
	if err := transactionExists(&transaction, id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Handle incoming JSON data.
	if err := handleTransactionData(w, r, &transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Save the updated transaction.
	db.DB.Save(&transaction)
	// Return the new transaction.
	json.NewEncoder(w).Encode(transaction)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if id is an integer.
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, "Id is invalid", http.StatusBadRequest)
		return
	}

	db.DB.Delete(&db.Transaction{}, id)

	fmt.Println(id)
}

func handleTransactionData(
	w http.ResponseWriter,
	r *http.Request,
	transaction *db.Transaction,
) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Set the incoming data to the transaction variable.
	if err := decoder.Decode(&transaction); err != nil {
		return err
	}

	if transaction.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if transaction.Amount == 0 {
		return errors.New("Amount cannot be empty")
	}

	return nil
}

func transactionExists(transaction *db.Transaction, id string) error {
	if err := db.DB.First(&transaction, id).Error; err != nil {
		return errors.New("Id is invalid")
	}

	return nil
}
