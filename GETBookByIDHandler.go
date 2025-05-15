package bookservice

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Get a book by id
func GetBookIDHandler(w http.ResponseWriter, r *http.Request) {
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
