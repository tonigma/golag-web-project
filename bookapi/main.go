package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Book struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    Title      string `json:"title"`
    ISBN       string `json:"isbn"`
    Author     string `json:"author"`
    Year       int    `json:"year"`
}

var db *gorm.DB

func initDB() {
    var err error
    db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&Book{})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    var books []Book
    db.Find(&books)
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var book Book
    db.First(&book, params["id"])
    json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    json.NewDecoder(r.Body).Decode(&book)
    db.Create(&book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var book Book
    db.First(&book, params["id"])
    json.NewDecoder(r.Body).Decode(&book)
    db.Save(&book)
    json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var book Book
    db.Delete(&book, params["id"])
    json.NewEncoder(w).Encode("Book deleted")
}

func main() {
    initDB()

    r := mux.NewRouter()

    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

    http.ListenAndServe(":8000", r)
}
