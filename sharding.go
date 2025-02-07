package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbs = make(map[int]*sql.DB)

func init() {
	// Connect to all shards
	dbs[0] = connectToDB("root:12345@tcp(localhost:3306)/one")
	dbs[1] = connectToDB("root:12345@tcp(localhost:3306)/two")
	dbs[2] = connectToDB("root:12345@tcp(localhost:3306)/three")
}

func connectToDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func insertUser(userID int, name, email string) {
	shard := userID % 3
	query := "INSERT INTO users (user_id, name, email) VALUES (?, ?, ?)"
	_, err := dbs[shard].Exec(query, userID, name, email)
	if err != nil {
		log.Printf("Error inserting into shard %d: %v", shard, err)
	} else {
		fmt.Printf("Inserted user %d into shard %d\n", userID, shard)
	}
}

func main() {
	insertUser(1001, "Alice", "alice@example.com")
	insertUser(1002, "Bob", "bob@example.com")
	insertUser(1003, "Charlie", "charlie@example.com")
}
