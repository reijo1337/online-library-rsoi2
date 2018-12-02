package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reijo1337/online-library-rsoi2/clients"
)

var (
	BooksPartClient   clients.BooksPartInterface
	ReadersPartClient clients.ReadersPartInterface
	ArrearsPartClient clients.ArrearsPartInterface
)

func init() {
	log.Println("Gateway: Initing...")
	bpc, err := clients.NewBooksPart()
	if err != nil {
		log.Println("Initing books part error")
		panic(err)
	}

	rpc, err := clients.NewReadersPart()
	if err != nil {
		log.Println("Initing readers part error")
		panic(err)
	}

	apc, err := clients.NewArrearsPart()
	if err != nil {
		log.Println("Initing arrears part error")
		panic(err)
	}

	BooksPartClient = bpc
	ReadersPartClient = rpc
	ArrearsPartClient = apc
}

// 5. Должен быть хотя бы один запрос, требующий данных с нескольких сервисов.
// Т.е. с Gateway выполняется запрос данных с нескольких сервисов (двух и более) и их агрегация.
// 7. При получении списка данных предусмотреть пагинацию.
// Запрос записанных на юзера книг по имени
func getUserArrears(c *gin.Context) {
	name := c.Query("name")
	pageSize := c.DefaultQuery("size", "5")
	pageNumber := c.DefaultQuery("page", "1")
	log.Println("Gateway: New request for arrears. Reader name", name, ", page:", pageNumber, ", size:", pageSize)
	log.Println("Gateway: Converting strings to numbers")
	page, err := strconv.ParseInt(pageNumber, 10, 32)
	if err != nil {
		log.Println("Gateway: Error in parsing page number:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Некорректно задан номер страницы",
			},
		)
		return
	}
	pageInt := int32(page)
	size, err := strconv.ParseInt(pageSize, 10, 32)
	if err != nil {
		log.Println("Gateway: Error in parsing page size:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Некорректно задан размер страницы",
			},
		)
		return
	}
	sizeInt := int32(size)

	log.Println("Gateway: Getting reader from remote service")
	reader, err := ReadersPartClient.GetReaderByName(name)
	if err != nil {
		log.Println("Gateway: Error while getting reader by name:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Нет возможности узнать, записан ли " + name + " в библиотеке",
			},
		)
		return
	}

	log.Println("Gateway: Getting arrears from remote service", reader.ID, pageInt, sizeInt)
	arrears, err := ArrearsPartClient.GetArrearsPaging(reader.ID, pageInt, sizeInt)
	if err != nil {
		log.Println("Gateway: Error while getting arrears:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Нет возможности получить информацию про книги, записанные на " + name,
			},
		)
		return
	}
	ret := []gin.H{}
	for _, ar := range arrears {
		book, err := BooksPartClient.GetBookByID(ar.BookID)
		if err != nil {
			log.Println("Gateway: Error while getting book by ID:", err.Error())
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "Нет возможности получить информацию про книги, записанные на " + name,
				},
			)
			return
		}
		ret = append(ret, gin.H{
			"id":          ar.ID,
			"reader_id":   ar.ReaderID,
			"book_id":     ar.BookID,
			"book_name":   book.Name,
			"book_author": book.Author.Name,
			"start":       ar.Start,
			"end":         ar.End,
		})
	}

	log.Println("Gateway: Request processed succesfully")
	c.JSON(http.StatusOK, ret)
}

// 6. Должно быть минимум два запроса, выполняющие обновление данных на нескольких сервисах в рамках одной операции.
// Запись книги на пользователя
func newArear(c *gin.Context) {
	log.Println("Gateway: New request for making new arrear")
	var req clients.NewReaderWithArrearRequestBody

	if err := c.BindJSON(&req); err != nil {
		log.Println("Gateway: Can't parse request body:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Пробелмы с обработкой запроса",
			},
		)
		return
	}
	log.Println("Gateway: Reader name:", req.ReaderName, ", Book ID:", req.BookID)
	log.Println("Gateway: Getting book by ID")
	book, err := BooksPartClient.GetBookByID(req.BookID)
	if err != nil {
		log.Println("Gateway: Can't recieve book:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Проблема с записью данной книги, возможно ее уже забрали.",
			},
		)
		return
	}
	log.Println("Gateway: Getting reader by name")
	reader, err := ReadersPartClient.GetReaderByName(req.ReaderName)
	if err != nil {
		log.Println("Gateway: Can't recieve reader:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Нет возможности узнать, записан ли " + req.ReaderName + " в библиотеке",
			},
		)
		return
	}
	log.Println("Gateway: Changing arrear book status to NOT FREE")
	err = BooksPartClient.ChangeBookStatusByID(req.BookID, false)
	if err != nil {
		log.Println("Gateway: Can't change book status:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Проблемы с резервированием книги",
			},
		)
		return
	}
	log.Println("Gateway: Making new arrear")
	arrear, err := ArrearsPartClient.NewArrear(reader.ID, req.BookID)
	if err != nil {
		log.Println("Gateway: Can't register new arrear:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Проблемы с записью книги. Попробуйте повторить запрос позже",
			},
		)
		return
	}
	log.Println("Gateway: Request processed succesfully")
	c.JSON(
		200,
		gin.H{
			"id":          arrear.ID,
			"reader_id":   arrear.ReaderID,
			"book_id":     arrear.BookID,
			"start":       arrear.Start,
			"end":         arrear.End,
			"book_name":   book.Name,
			"book_author": book.Author.Name,
		},
	)
}

// Возврат книги
func closeArrear(c *gin.Context) {
	arrearIDString := c.Query("id")
	log.Println("Gateway: New request for closing arrear with ID", arrearIDString)
	arrearID64, err := strconv.ParseInt(arrearIDString, 10, 32)
	if err != nil {
		log.Println("Gateway: Error in parsing arrear ID number:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Некорректно задан номер записи",
			},
		)
		return
	}
	arrearID := int32(arrearID64)

	log.Println("Gateway: Checking arrear for existanse by ID")
	arrear, err := ArrearsPartClient.GetArrearByID(arrearID)
	if err != nil {
		log.Println("Gateway: Error in getting arrear:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Нет возможности получить информацию о записи",
			},
		)
		return
	}
	log.Println("Gateway: Closing arrear by ID")
	err = ArrearsPartClient.CloseArrearByID(arrearID)
	if err != nil {
		log.Println("Gateway: Error in closing arrear:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Проблема с закрытием записи",
			},
		)
		return
	}

	log.Println("Gateway: Changing status of book from arrear to FREE")
	err = BooksPartClient.ChangeBookStatusByID(arrear.BookID, true)
	if err != nil {
		log.Println("Gateway: Can't change book status:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Проблема с закрытием записи",
			},
		)
		return
	}
	log.Println("Gateway: Request processed successfully")
	c.JSON(
		200,
		gin.H{
			"ok": "ok",
		},
	)
}

// Список доступных книг
func freeBooks(c *gin.Context) {
	log.Println("Gateway: New request for free books")
	log.Println("Gateway: Getting books from remote service")
	books, err := BooksPartClient.GetFreeBooks()
	if err != nil {
		log.Println("Gateway: Error while getting free books:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Проблемы с получением списка доступных книг",
			},
		)
		return
	}
	ret := []gin.H{}
	for _, bk := range books {
		ret = append(ret, gin.H{
			"id":        bk.ID,
			"name":      bk.Name,
			"author":    bk.Author.Name,
			"author_id": bk.Author.ID,
		})
	}

	log.Println("Gateway: Request processed succesfully")
	c.JSON(http.StatusOK, ret)
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getUserArrears", getUserArrears)
	r.POST("/arrear", newArear)
	r.DELETE("/arrear", closeArrear)
	r.GET("/freeBooks", freeBooks)
	r.OPTIONS("/arrear", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
	r.OPTIONS("/freeBooks", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := SetUpRouter()

	r.Run(":" + port)
}
