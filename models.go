package main

// Writer описывает модель писателя, чьи книги есть в библиотеке
type Writer struct {
	ID   int32
	Name string
}

// Book описывает модель книги, которая есть в библиотеке
type Book struct {
	ID     int32
	Name   string
	Author *Writer
	Free   bool
}

// Reader описывает модель читателя, который может брать книги в библиотеке
type Reader struct {
	ID   int32
	Name string
}

type Arrear struct {
	ID       int32
	readerID int32
	bookID   int32
	start    string
	end      string
}

type NewReaderWithArrearRequestBody struct {
	ReaderName string `json:"reader"`
	BookID     int32  `json:"book"`
}
