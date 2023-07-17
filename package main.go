package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Phone  string `json:"phone"`
}

func main() {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "user=nekiarie password=Kiariew@njiiri94 dbname=postgresql sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Execute a query to retrieve data from the user_table
	rows, err := db.Query("SELECT user_id, name, age, phone FROM public.user_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Create a slice to store the retrieved users
	users := []User{}

	// Iterate over the result set and populate the users slice
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Phone)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	// Convert the users slice to JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}
