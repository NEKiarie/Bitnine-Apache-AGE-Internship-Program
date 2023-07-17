package pgdriver

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
)

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Phone  string `json:"phone"`
}

type UserTable struct {
	db *sql.DB
}

func NewUserTable(db *sql.DB) *UserTable {
	return &UserTable{
		db: db,
	}
}

func (ut *UserTable) GetUsers() ([]User, error) {
	rows, err := ut.db.Query("SELECT user_id, name, age, phone FROM public.user_table")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ut *UserTable) GetUsersJSON() ([]byte, error) {
	users, err := ut.GetUsers()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
