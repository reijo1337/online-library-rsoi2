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
}

// Reader описывает модель читателя, который может брать книги в библиотеке
type Reader struct {
	ID   int32
	Name string
}
