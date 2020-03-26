package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// users slice holds a list of users
var books = []*Book{}

// Set up some initial data
func init() {
	fmt.Println("Setting up initial books")
	var author = NewAuthor("Ann", "Ballard")
	var publisher = NewPublisher("Tech Books Publishing Ltd.", "10 High Street")
	books = append(books, NewBook(123, "Go Servers", author, "technical", publisher, 12.99))
	books = append(books, NewBook(123, "Go Games", author, "technical", publisher, 12.99))
}

// --------------------------
// Sales interface
// --------------------------

// Sales an interface for items that can be put on sale
type Sales interface {
	setSaleDiscount(discount float64)
	calculateSalePrice() float64
}

// --------------------------
// Product struct
// --------------------------

// Product - represents all products sold by the shop
type Product struct {
	Title    string
	Price    float64
	Discount float64
}

func (p *Product) setSaleDiscount(discount float64) {
	p.Discount = discount
}

func (p *Product) calculateSalePrice() float64 {
	discountAmount := p.Price * p.Discount
	return p.Price - discountAmount
}

// --------------------------
// Book
// --------------------------

// Book structure
type Book struct {
	ISBN int
	Product
	Author    *Author
	Genre     string
	Publisher *Publisher
}

// NewBook constructor function for Books
func NewBook(isbn int, title string, author *Author, genre string, publisher *Publisher, price float64) *Book {
	book := Book{isbn, Product{title, price, 0.0}, author, genre, publisher}
	return &book
}

func (b *Book) prettyPrint() {
	fmt.Print("Book(")
	fmt.Printf("%d, ", b.ISBN)
	fmt.Printf("%s, ", b.Title)
	fmt.Printf("Author (%s, %s), ", b.Author.FirstName, b.Author.Surname)
	fmt.Printf("%s, ", b.Genre)
	fmt.Printf(", Publisher(%s, %s), ", b.Publisher.Name, b.Publisher.Address)
	fmt.Printf("%0.2f, ", b.Price)
	fmt.Printf("%0.2f)", b.Discount)
	fmt.Println()
}

// Stringer interface
func (b *Book) String() string {
	return fmt.Sprintf("Book(%d, %s, %v, %s, %v, %02.f, %0.2f)", b.ISBN, b.Title, b.Author, b.Genre, b.Publisher, b.Price, b.Discount)
}

// --------------------------
// End of Book
// --------------------------

// Author is an author
type Author struct {
	FirstName string
	Surname   string
}

// NewAuthor constructor function for Authors
func NewAuthor(firstName string, surname string) *Author {
	author := Author{firstName, surname}
	return &author
}

// Stringer interface
func (a *Author) String() string {
	return fmt.Sprintf("Author(%s, %s)", a.FirstName, a.Surname)
}

// -------------------------

// Publisher of books
type Publisher struct {
	Name    string
	Address string
}

// NewPublisher constructor function for Publishers
func NewPublisher(name string, address string) *Publisher {
	publisher := Publisher{name, address}
	return &publisher
}

// Stringer interface
func (p *Publisher) String() string {
	return fmt.Sprintf("Publisher(%s, %s)", p.Name, p.Address)
}

func getAllBooks(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	booksJSON, err := json.Marshal(books)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(booksJSON)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// Handle path parameters
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	isbn := -1
	var err error
	if val, ok := pathParams["isbn"]; ok {
		isbn, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a isbn"}`))
			return
		}
	}
	var author = NewAuthor("Pete", "Smith")
	var publisher = NewPublisher("Tech Books Publishing Ltd.", "10 High Street")
	var book = NewBook(isbn, "Go Unleashed", author, "Technical", publisher, 15.99)

	fmt.Println("book:", book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&book); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "problem generating user info"}`))
	}
}

func notFound(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(`{"message": "not found"}`))
}

func main() {
	fmt.Println("Starting Server")

	const urlPath = "/api/"
	router := mux.NewRouter()
	// Create a subrouter - so don;t have to repeat path info
	api := router.PathPrefix(urlPath).Subrouter()
	api.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)
	api.HandleFunc("/books/{isbn}", getBook).Methods(http.MethodGet)

	api.HandleFunc("", notFound)
	fmt.Println("Server Available - see")
	fmt.Println("\thttp://localhost:8080/api/books")
	fmt.Println("\thttp://localhost:8080/api/books/123")
	err := http.ListenAndServe(":8080", api)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
