How to start the application:

1. Install the necessary dependencies (gorilla/mux and gorm with sqlite driver):

    go get -u github.com/gorilla/mux
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/sqlite

2. Run the application using the following command:

    go run main.go

3. Test the Endpoints:

    Get all books: GET http://localhost:8000/books
    Get a single book: GET http://localhost:8000/books/{id}
    Create a new book: POST http://localhost:8000/books

Format of the expected data is: 
{
    "title": "Harry Poter",
    "isbn": "123-456-789",
    "author": "J. K. Rowling",
    "year": 1997
}

To update the data use:

    Update a book: PUT http://localhost:8000/books/{id}

as a body set:
{
    "title": "Harry Poter 2",
    "isbn": "123-456-789",
    "author": "J. K. Rowling",
    "year": 1998
}

to delete it just use:


  Delete a book: DELETE http://localhost:8000/books/{id}


