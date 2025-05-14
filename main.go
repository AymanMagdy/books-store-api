package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import the driver anonymously

	"github.com/gorilla/mux"
)

type Book struct {
	ID int   	`json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Description string `json:"description"`
}

const database string = "./data/mydb.sqlite"

func createTable () bool {
	const create string = `
		CREATE TABLE IF NOT EXISTS books (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		author TEXT NOT NULL,
		description TEXT NOT NULL 
	);`

	db, err := sql.Open("sqlite3", database)

	if err != nil {
		return false
	}

	if _, err := db.Exec(create); err != nil {
		return false
	}
	
	db.Close()
	return true

}

// Returns a list of books
func getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed to get all books", http.StatusMethodNotAllowed)
		return
	}

	db, err := sql.Open("sqlite3", database)

	if err != nil {
		return 
	}

	query := `SELECT * FROM books;`

	rows, err := db.Query(query)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}

	if _, err := db.Exec(query); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} 
	
	var books []Book

	for rows.Next(){
		var book Book

		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		books = append(books, book)
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db.Close()
	w.WriteHeader(http.StatusOK)
}

// Delete a book...
func deleteBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is allowed to delete a book", http.StatusMethodNotAllowed)
		return
	}

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return 
	}

	vars := mux.Vars(r)
	id := vars["id"]

	query := `DELETE FROM books WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if result != nil {
		affectedRows, err := result.RowsAffected()
		if err != nil{
			println(("Error deleting rows"))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if affectedRows == 0{
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		db.Close()
		log.Println("Deleted rows: ", affectedRows)
		w.WriteHeader(http.StatusOK)

	}

}

// Get a book by id
func getBookIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed to get a book", http.StatusMethodNotAllowed)
		return
	}
	db, err := sql.Open("sqlite3", database)

	vars := mux.Vars(r)
	id := vars["id"]

	if err != nil {
		return 
	}

	query := `SELECT * FROM books WHERE id = ?`

	row := db.QueryRow(query, id)

	if _, err := db.Exec(query, id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} 

	var book Book
	if err := row.Scan(&book.ID, &book.Name, &book.Author, &book.Description); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	bookResult := Book {
		ID:  book.ID,
		Name: book.Name,
		Author: book.Author,
		Description: book.Description,
	}
	
	if err := json.NewEncoder(w).Encode(bookResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	db.Close()
	
}

// Update a book
func updateBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is allowed to update", http.StatusMethodNotAllowed)
		return
	}
	
	db, err := sql.Open("sqlite3", database)

	if err != nil{
		log.Print("Error connecting to DB...")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the book exists
	row := db.QueryRow(`SELECT * FROM books WHERE id = ?`, id)

	if _, err := db.Exec(`SELECT * FROM books WHERE id = ?`, id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} 

	var testBoook Book
	if err := row.Scan(&testBoook.ID, &testBoook.Name, &testBoook.Author, &testBoook.Description); err != nil {
		log.Printf("Book with id: %s not found", id)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var book Book

	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := `UPDATE books SET name=?, author=?, description=? WHERE id = ?`

	db.QueryRow(query, book.Name, book.Author, book.Description, id)

	if _, err := db.Exec(query, book.Name, book.Author, book.Description, id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} 

	log.Println("Updated Book with id: ", id)
	w.WriteHeader(http.StatusOK)

}

// Create a new book
func createBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Only PUT method is allowed to update", http.StatusMethodNotAllowed)
		return
	}
	
	db, err := sql.Open("sqlite3", database)

	if err != nil{
		log.Print("Error connecting to DB...")
	}

	var book Book

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil{
		log.Println("Error reading the JSON")
	}

	query := `INSERT INTO books(name, author, description) VALUES(?, ?, ?)`

	row, err := db.Query(query, book.Name, book.Author, book.Description)

	if err != nil{
		log.Println("Error inserting a new book")
	}

	if _, err := db.Exec(query, book.Name, book.Author, book.Description); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Print(row)
	w.WriteHeader(http.StatusOK)

}

func main() {

	db := createTable()
	if db == false{
		log.Println("Error creating the table.")
	}
	router := mux.NewRouter()
	router.HandleFunc("/books", getAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/{id}", getBookIDHandler).Methods("GET")
	router.HandleFunc("/books/{id}", deleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books/{id}", updateBookHandler).Methods("PUT")
	router.HandleFunc("/books", createBookHandler).Methods("POST") // here...


	fmt.Println("Server is running at http://localhost:8080")

	http.ListenAndServe(":8080", router)
}
