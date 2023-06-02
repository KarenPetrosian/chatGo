package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
}

func validateUsername(user User) bool {
	if len(user.Username) > 0 {
		return true
	}
	return false
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	writeUser()
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer request.Body.Close()
	jsonResponse, err := json.Marshal(body)
	if err != nil {
		http.Error(response, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func writeUser() {
	db, err := sql.Open("mysql", "username:password@tcp(http://127.0.0.1:3306)/user")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert a new user
	user := User{
		Username: "john_doe",
	}
	err = insertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User inserted successfully!")
}

func insertUser(db *sql.DB, user User) error {
	query := "INSERT INTO users (username, email) VALUES (?, ?)"
	_, err := db.Exec(query, user.Username)
	return err
}
