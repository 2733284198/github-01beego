package models

type Book struct {
	Title   string
	Author  string
	Subject string
	Bookid  int
}

type Books struct {
	Books []Book
}
