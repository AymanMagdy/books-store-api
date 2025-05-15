package bookservice

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Returns a list of books
func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
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