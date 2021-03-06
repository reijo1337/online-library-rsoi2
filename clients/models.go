package clients

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
	ReaderID int32
	BookID   int32
	Start    string
	End      string
}

type NewReaderWithArrearRequestBody struct {
	ReaderName string `json:"reader"`
	BookID     int32  `json:"book"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}
