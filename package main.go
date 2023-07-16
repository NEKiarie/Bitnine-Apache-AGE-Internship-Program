package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Phone  string `json:"phone"`
}

func main() {
	db, err := sql.Open("postgres", "user=nekiarie password=Kiariew@njiiri94 dbname=postgresql sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM public.user_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
