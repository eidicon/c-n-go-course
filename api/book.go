package api

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Book type
type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// stubdata
var books = map[string]Book{
	"4041235780": Book{Title: "The Hitchhiker's guide to the Galaxy", Author: "Douglas Adams", ISBN: "4041235780" },
	"0123456789": Book{Title: "Cloud native go", Author: "M.-L. Reimer", ISBN: "0123456789" },
}

// ToJSON deserialize Book sruct to a json
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return ToJSON
}

//FromJSON serialize json data to the Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}

	return book
}

//BooksHandleFunc Controller for route
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := GetAllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := 	ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/v1/books/" + isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method!"))
	}
}

// BookHandleFunc to be used as http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/v1/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

//GetAllBooks returns list of books 
func GetAllBooks() []Book {	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}