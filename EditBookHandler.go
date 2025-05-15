package bookservice

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Update a book
func UpdateBookHandler(w http.ResponseWriter, r *http.Request){
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