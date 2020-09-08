package router

import (
	"github.com/gorilla/mux"
	"github.com/kangbojk/library_api/entity"
)

func NewRouter(repo book.Repository) *mux.Router {
	m := mux.NewRouter()

	m.HandleFunc("/api/books", listBooks(repo)).Methods("GET")
	m.HandleFunc("/api/books", createBook(repo)).Methods("POST")
	m.HandleFunc("/api/books/{id}", getBook(repo)).Methods("GET")
	m.HandleFunc("/api/books/{id}", updateBook(repo)).Methods("PUT")
	m.HandleFunc("/api/books/{id}", deleteBook(repo)).Methods("DELETE")

	return m
}
