package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zohaibsoomro/book-server-mongodb/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	books, err := model.GetAllBooksFromDb()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	s, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(s)
}

func GetBookById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id := p.ByName("id")
	iod, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Book id not formatted.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	b, err := model.GetBookWithIdFromDB(iod)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book id not found!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(s)
}

func CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	b := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Request not have valid body.\n"))
		w.Write([]byte(err.Error()))
		return
	}

	if err := b.CreateBookInDB(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Book creation failed.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.Write([]byte("Book created\n"))
	w.Write(s)
	w.WriteHeader(http.StatusCreated)

}

func UpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id := p.ByName("id")
	iod, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Book id not formatted.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	b, err := model.GetBookWithIdFromDB(iod)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id not found!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	book := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Request not have valid body.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	if b.Author != book.Author {
		b.Author = book.Author
	}
	if b.Name != book.Name {
		b.Name = book.Name
	}
	if b.PublishDate != book.PublishDate {
		b.PublishDate = book.PublishDate
	}
	if err = b.UpdateBookInDb(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Update failed!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Updated\n"))
	w.Write(s)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id := p.ByName("id")
	iod, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Book id not formatted.\n"))
		w.Write([]byte(err.Error()))
		return
	}
	b, err := model.GetBookWithIdFromDB(iod)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id not found!\n"))
		return
	}
	book, err := b.DeleteBookWithIdFromDb()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Deletion failed!\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted book:\n"))
	w.Write(s)
}
