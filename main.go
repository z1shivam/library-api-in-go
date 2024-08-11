package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 3},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 5},
	{ID: "3", Title: "1984", Author: "George Orwell", Quantity: 2},
	{ID: "4", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 7},
	{ID: "5", Title: "Moby-Dick", Author: "Herman Melville", Quantity: 4},
	{ID: "6", Title: "Pride and Prejudice", Author: "Jane Austen", Quantity: 6},
	{ID: "7", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 3},
	{ID: "8", Title: "To the Lighthouse", Author: "Virginia Woolf", Quantity: 1},
	{ID: "9", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 8},
	{ID: "10", Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "book cannot be created"})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkinBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id parameter."})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusFound, gin.H{"message": "Book not available!"})
		return
	}
	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id parameter."})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity++
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)

	router.POST("/books", createBook)
	router.POST("/checkin", checkinBook)
	router.POST("checkout", checkoutBook)

	router.Run("localhost:8080") // this run the server!
}
