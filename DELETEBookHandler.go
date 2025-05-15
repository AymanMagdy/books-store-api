package bookservice

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Delete a book...
func DeleteBookHandler(w http.ResponseWriter, r *http.Request){
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