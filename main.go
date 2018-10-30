package main

import (
	"log"
	"net/http"
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
	log.Println("Gateway: Initing...")
	bpc, err := NewBooksPart()
	if err != nil {
		log.Println("Initing books part error")
		panic(err)
	}

	rpc, err := NewReadersPart()
	if err != nil {
		log.Println("Initing readers part error")
		panic(err)
	}

	apc, err := NewArrearsPart()
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
	pageSize := c.DefaultQuery("size", "10")
	pageNumber := c.DefaultQuery("page", "1")
	log.Println("Gateway: New request for arrears. Reader name", name, ", page:", pageNumber, ", size:", pageSize)
	log.Println("Gateway: Converting strings to numbers")
	page, err := strconv.ParseInt(pageNumber, 10, 32)
	if err != nil {
		log.Println("Gateway: Error in parsing page number:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}
	pageInt := int32(page)
	size, err := strconv.ParseInt(pageSize, 10, 32)
	if err != nil {
		log.Println("Gateway: Error in parsing page size:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	sizeInt := int32(size)

	log.Println("Gateway: Getting reader from remote service")
	reader, err := ReadersPartClient.getReaderByName(name)
	if err != nil {
		log.Println("Gateway: Error while getting reader by name:", err.Error())
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	log.Println("Gateway: Getting arrears from remote service")
	arrears, err := ArrearsPartClient.getArrearsPaging(reader.ID, pageInt, sizeInt)
	if err != nil {
		log.Println("Gateway: Error while getting arrears:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	ret := gin.H{}
	for i, ar := range arrears {
		book, err := BooksPartClient.getBookByID(ar.bookID)
		if err != nil {
			log.Println("Gateway: Error while getting book by ID:", err.Error())
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"error": err.Error(),
				},
			)
			return
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

	log.Println("Gateway: Request processed succesfully")
	c.JSON(http.StatusOK, ret)
}

// 6. Должно быть минимум два запроса, выполняющие обновление данных на нескольких сервисах в рамках одной операции.
// Запись книги на пользователя
func newArear(c *gin.Context) {
	log.Println("Gateway: New request for making new arrear")
	var req NewReaderWithArrearRequestBody
	if err := c.ShouldBind(&req); err != nil {
		log.Println("Gateway: Can't parse request body:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Reader name:", req.ReaderName, ", Book ID:", req.BookID)
	log.Println("Gateway: Getting book by ID")
	if _, err := BooksPartClient.getBookByID(req.BookID); err != nil {
		log.Println("Gateway: Can't recieve book:", err.Error())
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Getting reader by name")
	reader, err := ReadersPartClient.getReaderByName(req.ReaderName)
	if err != nil {
		log.Println("Gateway: Can't recieve reader:", err.Error())
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Changing arrear book status to NOT FREE")
	err = BooksPartClient.changeBookStatusByID(req.BookID, false)
	if err != nil {
		log.Println("Gateway: Can't change book status:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Making new arrear")
	arrear, err := ArrearsPartClient.newArrear(reader.ID, req.BookID)
	if err != nil {
		log.Println("Gateway: Can't register new arrear:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Request processed succesfully")
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
				"error": err.Error(),
			},
		)
		return
	}
	arrearID := int32(arrearID64)

	log.Println("Gateway: Checking arrear for existanse by ID")
	arrear, err := ArrearsPartClient.getArrearByID(arrearID)
	if err != nil {
		log.Println("Gateway: Error in getting arrear:", err.Error())
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Closing arrear by ID")
	err = ArrearsPartClient.closeArrearByID(arrearID)
	if err != nil {
		log.Println("Gateway: Error in closing arrear:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	log.Println("Gateway: Changing status of book from arrear to FREE")
	err = BooksPartClient.changeBookStatusByID(arrear.bookID, true)
	if err != nil {
		log.Println("Gateway: Can't change book status:", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	log.Println("Gateway: Request processed succesfully")
	c.JSON(
		200,
		gin.H{
			"ok": "ok",
		},
	)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := gin.Default()

	r.GET("/getUserArrears", getUserArrears)
	r.POST("/newArear", newArear)
	r.DELETE("/closeArrear", closeArrear)

	r.Run(":" + port)
}
