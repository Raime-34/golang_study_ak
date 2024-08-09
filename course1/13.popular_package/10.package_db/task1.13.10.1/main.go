package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserTable() error {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS users (" +
			"id INTEGER PRIMARY KEY, " +
			"name TEXT, " +
			"age INTEGER" +
			")",
	)

	return err
}

func InsertUser(user User) error {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	p, _ := db.Prepare(
		"INSERT INTO " +
			"users " +
			"VALUES(" +
			"?," +
			"?," +
			"?" +
			")",
	)

	defer p.Close()

	_, err := p.Exec(user.ID, user.Name, user.Age)

	return err
}

func SelectUser(id int) (User, error) {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	rows, _ := db.Query(
		"SELECT * "+
			"FROM users "+
			"WHERE id = ?",
		id,
	)
	defer rows.Close()

	if rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.Age)

		return user, nil
	}

	return User{}, fmt.Errorf("user not found")
}

func UpdateUser(user User) error {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	result, err := db.Exec(
		"UPDATE users SET name = ?, age = ? WHERE id = ?",
		user.Name, user.Age, user.ID,
	)

	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("user not found")
	}

	return err
}

func DeleteUser(id int) error {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	result, err := db.Exec(
		"DELETE FROM users WHERE ?",
		id,
	)

	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("user not found")
	}

	return err
}
