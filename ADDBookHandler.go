package bookservice

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Create a new book
func CreateBookHandler(w http.ResponseWriter, r *http.Request){
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