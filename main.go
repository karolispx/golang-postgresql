package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "admin"
	DB_PASSWORD = "root"
	DB_NAME     = "database"
)

func main() {
	// DB connection
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	defer db.Close()

	printMessage("Delete books from DB...")
	deleteBooks(db)

	printMessage("Insert books into DB...")
	insertBooks(db)

	printMessage("Getting books...")
	getBooks(db)
}

// Insert books
func insertBooks(db *sql.DB) {
	// Insert several books
	for i := 0; i < 10; i++ {
		printMessage("Inserting book into DB")

		// Convert i to string - for demo purposes only
		var thisBookID_string = strconv.Itoa(i)

		var bookID = thisBookID_string

		var bookName = thisBookID_string + "_name"

		fmt.Println("Inserting new book with ID: " + bookID + " and name: " + bookName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO books(bookID, bookName) VALUES($1, $2) returning id;", bookID, bookName).Scan(&lastInsertID)

		checkErr(err)
	}
}

// Get books
func getBooks(db *sql.DB) {
	// Get all books from books table that don't have bookID = "1"
	rows, err := db.Query("SELECT * FROM books where bookID <> $1", "1")

	checkErr(err)

	// Foreach book
	for rows.Next() {
		var id int
		var bookID string
		var bookName string

		err = rows.Scan(&id, &bookID, &bookName)

		checkErr(err)

		fmt.Println("Book ID: " + bookID + ", book name: " + bookName)
	}
}

// Delete books
func deleteBooks(db *sql.DB) {
	printMessage("Deleting all books...")

	_, err := db.Exec("DELETE FROM books")
	checkErr(err)

	printMessage("All books have been deleted successfully!")
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
