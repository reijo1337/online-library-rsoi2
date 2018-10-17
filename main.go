package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	BooksPartClient   *BooksPart
	ReadersPartClient *ReadersPart
)

func init() {
	bpc, err := NewBooksPart()
	if err != nil {
		panic(err)
	}

	rpc, err := NewReadersPart()
	if err != nil {
		panic(err)
	}

	BooksPartClient = bpc
	ReadersPartClient = rpc
}

// 5. Должен быть хотя бы один запрос, требующий данных с нескольких сервисов.
// Т.е. с Gateway выполняется запрос данных с нескольких сервисов (двух и более) и их агрегация.
// 7. При получении списка данных предусмотреть пагинацию.
// Запрос записанных на юзера книг по ID
func getUserArrears(c *gin.Context) {
	ID := c.Param("id")
	pageSize := c.DefaultQuery("size", "10")
	pageNumber := c.DefaultQuery("page", "1")
	i, err := strconv.ParseInt(ID, 10, 32)
	if err != nil {
		panic(err)
	}
	result := int32(i)

	reader, err := ReadersPartClient.getReaderByID(result)
	if err != nil {
		panic(err)
	}
}

// 6. Должно быть минимум два запроса, выполняющие обновление данных на нескольких сервисах в рамках одной операции.
// Регистрация нового пользователя и запись на него новой книги
func newReaderWithArear(c *gin.Context) {

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

	r.GET("/getUserArrears/:id", getUserArrears)

	r.Run(":" + port)
}
