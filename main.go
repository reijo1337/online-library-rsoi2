package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// 5. Должен быть хотя бы один запрос, требующий данных с нескольких сервисов.
// Т.е. с Gateway выполняется запрос данных с нескольких сервисов (двух и более) и их агрегация.
// Запрос записанных на юзера книг по ID
func getUserArrears(c *gin.Context) {

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

	r.Run(":" + port)
}
