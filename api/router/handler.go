package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kangbojk/library_api/api/presenter"
	"github.com/kangbojk/library_api/entity"
)

func listBooks(repo book.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var data []*book.Book
		var err error
		title := r.URL.Query().Get("title")
		switch {
		case title == "":
			data, err = repo.List()
		default:
			data, err = repo.Search(title)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != book.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(book.ErrNotFound.Error()))
			return
		}
		var toJson []*presenter.Book
		for _, d := range data {
			toJson = append(toJson, &presenter.Book{
				ID:       d.ID,
				Title:    d.Title,
				Author:   d.Author,
				Pages:    d.Pages,
				Quantity: d.Quantity,
			})
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func createBook(repo book.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var input struct {
			Title    string `json:"title"`
			Author   string `json:"author"`
			Pages    int    `json:"pages"`
			Quantity int    `json:"quantity"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		b := &book.Book{
			ID:        uuid.New(),
			Title:     input.Title,
			Author:    input.Author,
			Pages:     input.Pages,
			Quantity:  input.Quantity,
			CreatedAt: time.Now(),
		}
		b.ID, err = repo.Create(b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		toJson := &presenter.Book{
			ID:       b.ID,
			Title:    b.Title,
			Author:   b.Author,
			Pages:    b.Pages,
			Quantity: b.Quantity,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func getBook(repo book.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := repo.Get(id)
		if err != nil && err != book.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(book.ErrNotFound.Error()))
			return
		}
		toJson := &presenter.Book{
			ID:       data.ID,
			Title:    data.Title,
			Author:   data.Author,
			Pages:    data.Pages,
			Quantity: data.Quantity,
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(book.ErrNotFound.Error()))
		}
	}
}

func updateBook(repo book.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(book.ErrNotFound.Error()))
			return
		}

		var input struct {
			Title    string `json:"title"`
			Author   string `json:"author"`
			Pages    int    `json:"pages"`
			Quantity int    `json:"quantity"`
		}
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		b, err := repo.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if input.Title != "" {
			b.Title = input.Title
		}

		if input.Author != "" {
			b.Author = input.Author
		}

		if input.Pages != 0 {
			b.Pages = input.Pages
		}

		b.Quantity = input.Quantity
		b.UpdatedAt = time.Now()

		err = repo.Update(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("Successfully update book \"" + b.Title + "\"\n"))
	}
}

func deleteBook(repo book.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(book.ErrNotFound.Error()))
			return
		}
		err = repo.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(book.ErrNotFound.Error()))
			return
		}

		w.Write([]byte("Successfully delete book \n"))
	}
}
