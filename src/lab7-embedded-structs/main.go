package main

import (
	"fmt"
)

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

// ------------------------

func handleSalesitems(sales Sales) {
	fmt.Println("\nCalculating the Sales Discount price")
	sales.setSaleDiscount(0.10)
	fmt.Printf("Sale price: %0.2f", sales.calculateSalePrice())
}

func main() {
	fmt.Println("Welcome to the Bookshop")

	var author = NewAuthor("Pete", "Smith")
	var publisher = NewPublisher("Tech Books Publishing Ltd.", "10 High Street")
	var book = NewBook(1, "Scala Unleashed", author, "Technical", publisher, 15.99)
	book.prettyPrint()

	fmt.Println(author)
	fmt.Println(publisher)
	fmt.Println(book)

	handleSalesitems(book)
}
