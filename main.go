package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	BooksPartClient   *BooksPart
	ReadersPartClient *ReadersPart
	ArrearsPartClient *ArrearsPart
)

func init() {
	bpc, err := NewBooksPart()
	if err != nil {
		fmt.Println("Initing books part error")
		panic(err)
	}

	rpc, err := NewReadersPart()
	if err != nil {
		fmt.Println("Initing readers part error")
		panic(err)
	}

	apc, err := NewArrearsPart()
	if err != nil {
		fmt.Println("Initing arrears part error")
		panic(err)
	}

	BooksPartClient = bpc
	ReadersPartClient = rpc
	ArrearsPartClient = apc
}

// 5. Должен быть хотя бы один запрос, требующий данных с нескольких сервисов.
// Т.е. с Gateway выполняется запрос данных с нескольких сервисов (двух и более) и их агрегация.
// 7. При получении списка данных предусмотреть пагинацию.
// Запрос записанных на юзера книг по ID
func getUserArrears(c *gin.Context) {
	name := c.Query("name")
	fmt.Println("name: ", name)
	pageSize := c.DefaultQuery("size", "10")
	pageNumber := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageNumber, 10, 32)
	if err != nil {
		fmt.Printf("Error in parsing page number")
		panic(err)
	}
	pageInt := int32(page)
	size, err := strconv.ParseInt(pageSize, 10, 32)
	if err != nil {
		fmt.Printf("Error in parsing page size")
		panic(err)
	}
	sizeInt := int32(size)

	reader, err := ReadersPartClient.getReaderByName(name)
	if err != nil {
		fmt.Printf("Error while getting reader by name")
		c.JSON(
			404,
			gin.H{
				"error": err.Error(),
			})
		c.Error(err)
	}

	arrears, err := ArrearsPartClient.getArrearsPaging(reader.ID, pageInt, sizeInt)
	if err != nil {
		panic(err)
	}
	ret := gin.H{}
	for i, ar := range arrears {
		book, err := BooksPartClient.getBookByID(ar.bookID)
		if err != nil {
			fmt.Printf("Error while getting reader by name")
			c.JSON(
				404,
				gin.H{
					"error": err.Error(),
				})
			c.Error(err)
		}
		ret[strconv.Itoa(i)] = gin.H{
			"id":          ar.ID,
			"reader_id":   ar.readerID,
			"book_id":     ar.bookID,
			"book_name":   book.Name,
			"book_author": book.Author.Name,
			"star":        ar.start,
			"end":         ar.end,
		}
	}

	c.JSON(200, ret)
}

// 6. Должно быть минимум два запроса, выполняющие обновление данных на нескольких сервисах в рамках одной операции.
// Регистрация нового пользователя и запись на него новой книги
func newReaderWithArear(c *gin.Context) {
	var req NewReaderWithArrearRequestBody
	if err := c.ShouldBind(&req); err != nil {
		panic(err)
	}
	fmt.Println(req.ReaderName, req.BookID)
	if _, err := BooksPartClient.getBookByID(req.BookID); err != nil {
		panic(err)
	}
	newReader, err := ReadersPartClient.registerReader(req.ReaderName)
	if err != nil {
		panic(err)
	}
	arrear, err := ArrearsPartClient.newArrear(newReader.ID, req.BookID)
	if err != nil {
		panic(err)
	}
	c.JSON(
		200,
		gin.H{
			"id":        arrear.ID,
			"reader_id": arrear.readerID,
			"book_id":   arrear.bookID,
			"start":     arrear.start,
			"end":       arrear.end,
		},
	)
}

// Потеря читателем книги
func registerBookLost(c *gin.Context) {

}

// Получение списка доступных книг
func getFreeBooks(c *gin.Context) {

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := gin.Default()

	r.GET("/getUserArrears", getUserArrears)
	r.POST("/newReaderWithArear", newReaderWithArear)

	r.Run(":" + port)
}
