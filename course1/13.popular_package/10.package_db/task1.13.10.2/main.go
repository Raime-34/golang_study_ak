package main

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const (
	SELECT = "select"
	INSERT = "insert"
	UPDATE = "update"
	DELETE = "delete"
)

func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	switch operation {
	case SELECT:
		return sq.Select("*").From(table).Where(sq.Eq{"id": user.ID}).ToSql()
	case INSERT:
		return sq.Insert(table).Columns("id", "name", "age").Values(user.ID, user.Name, user.Age).ToSql()
	case UPDATE:
		return sq.Update(table).SetMap(map[string]interface{}{"name": user.Name, "age": user.Age}).Where(sq.Eq{"id": user.ID}).ToSql()
	case DELETE:
		return sq.Delete(table).Where(sq.Eq{"id": user.ID}).ToSql()
	}

	return "", nil, fmt.Errorf("unhandled operation")
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

	query, args, _ := PrepareQuery(INSERT, "users", user)

	_, err := db.Exec(query, args...)

	return err
}

func SelectUser(id int) (User, error) {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	query, args, _ := PrepareQuery(SELECT, "users", User{ID: id})

	rows, _ := db.Query(query, args...)
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

	query, args, _ := PrepareQuery(UPDATE, "users", user)

	result, err := db.Exec(
		query,
		args...,
	)

	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("user not found")
	}

	return err
}

func DeleteUser(id int) error {
	db, _ := sql.Open("sqlite3", "user.db")
	defer db.Close()

	query, args, _ := PrepareQuery(DELETE, "users", User{ID: id})

	result, err := db.Exec(
		query,
		args...,
	)

	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("user not found")
	}

	return err
}
