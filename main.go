package bookservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := CreateTable()
	if db == false{
		log.Println("Error creating the table.")
	}
	router := mux.NewRouter()
	router.HandleFunc("/books", GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookIDHandler).Methods("GET")
	router.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books/{id}", UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books", CreateBookHandler).Methods("POST") // here...


	fmt.Println("Server is running at http://localhost:8080")

	http.ListenAndServe(":8080", router)
}
