package main

type ArrearResponse struct {
	ID         int32  `json:"id"`
	ReaderID   int32  `json:"reader_id"`
	BookID     int32  `json:"book_id"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"book_author"`
	Start      string `json:"start"`
	End        string `json:"end"`
}

type NewArrearResponse struct {
	ID         int32  `json:"id"`
	ReaderID   int32  `json:"reader_id"`
	BookID     int32  `json:"book_id"`
	Start      string `json:"start"`
	End        string `json:"end"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"book_author"`
}
