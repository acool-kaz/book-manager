package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/acool-kaz/book-manager/pkg/models"
	"github.com/acool-kaz/book-manager/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allBooks := models.GetAllBooks()
	json.NewEncoder(w).Encode(allBooks)
	// res, err := json.Marshal(allBooks)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	bookDetails, _ := models.GetBookById(id)
	json.NewEncoder(w).Encode(bookDetails)
	// res, err := json.Marshal(bookDetails)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBook := &models.Book{}
	if err := utils.ParseBody(r, newBook); err != nil {
		fmt.Fprint(w, err)
		return
	}
	b := newBook.CreateBook()
	json.NewEncoder(w).Encode(b)
	// res, err := json.Marshal(b)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	book := models.DeleteBook(id)
	json.NewEncoder(w).Encode(book)
	// res, err := json.Marshal(book)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	json.NewEncoder(w).Encode(bookDetails)
	// res, err := json.Marshal(bookDetails)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}
