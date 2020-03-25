package main

import (
	"fmt"
)

// Book structure
type Book struct {
	ISBN      int
	Title     string
	Author    *Author
	Genre     string
	Publisher *Publisher
	Price     float64
	Discount  float64
}

func NewBook(isbn int, title string, author *Author, genre string, publisher *Publisher, price float64) *Book {
	book := Book{isbn, title, author, genre, publisher, price, 0.0}
	return &book
}

func (b *Book) setSaleDiscount(discount float64) {
	b.Discount = discount
}

func (b *Book) calculateSalePrice() float64 {
	discountAmount := b.Price * b.Discount
	return b.Price - discountAmount
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
}

// Author is an author
type Author struct {
	FirstName string
	Surname   string
}

func NewAuthor(firstName string, surname string) *Author {
	author := Author{firstName, surname}
	return &author
}

// Publisher of books
type Publisher struct {
	Name    string
	Address string
}

func NewPublisher(name string, address string) *Publisher {
	publisher := Publisher{name, address}
	return &publisher
}

func main() {
	fmt.Println("Welcome to the Bookshop")

	var author = NewAuthor("Pete", "Smith")
	var publisher = NewPublisher("Tech Books Publishing Ltd.", "10 High Street")
	var book = NewBook(1, "Scala Unleashed", author, "Technical", publisher, 15.99)
	book.prettyPrint()

	fmt.Println("\nCalculating the Sales Discount price")
	book.setSaleDiscount(0.10)
	fmt.Printf("Sale price of book: %0.2f", book.calculateSalePrice())
}
