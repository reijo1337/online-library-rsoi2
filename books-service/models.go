package main

type Writer struct {
	ID   int32
	Name string
}

type Book struct {
	ID     int32
	Name   string
	Author *Writer
}
