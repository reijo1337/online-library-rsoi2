package main

type Writer struct {
	ID int32
	Name string
}

type Book {
	ID int32
	Name string
	Author *Writer
}