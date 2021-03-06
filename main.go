package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/reijo1337/online-library-rsoi2/clients"
)

var (
	BooksPartClient   clients.BooksPartInterface
	ReadersPartClient clients.ReadersPartInterface
	ArrearsPartClient clients.ArrearsPartInterface
	AuthPartClient    clients.AuthPartInterface
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

	aupc, err := clients.NewAuthPart()
	if err != nil {
		log.Println("Initing arrears part error")
		panic(err)
	}

	BooksPartClient = bpc
	ReadersPartClient = rpc
	ArrearsPartClient = apc
	AuthPartClient = aupc
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

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Gateway: New authorized request")
		tokenString := c.Query("access_token")
		if tokenString == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Unauthorized",
				},
			)
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// hmacSampleSecret := os.Getenv("SECRET")
			hmacSampleSecret := []byte("secc")
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Next()
		} else {
			log.Println("Gateway: Authorization failed: ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Unauthorized",
				},
			)
		}
	}
}

func Login(c *gin.Context) {
	log.Println("Gateway: New request for login")
	req := &clients.User{}
	if err := c.BindJSON(req); err != nil {
		log.Println("Gateway: Can't parse request body:", err.Error())
		log.Println(req)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Пробелмы с обработкой запроса",
			},
		)
	}

	tokens, err := AuthPartClient.GetToken(req)
	if err != nil {
		log.Println("Gateway: Login failed: ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "Неудачная попытка авторизации",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		tokens,
	)
}

func Refresh(c *gin.Context) {
	log.Println("Gateway: New request for refresh token")
	tokens, err := AuthPartClient.RefreshToken(c.Query("refresh_token"))
	if err != nil {
		log.Println("Gateway: Login failed: ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Unauthorized",
			},
		)
	}
	c.JSON(
		http.StatusOK,
		tokens,
	)
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	authorized := r.Group("/", AuthRequired())
	authorized.GET("/getUserArrears", getUserArrears)
	authorized.POST("/arrear", newArear)
	authorized.DELETE("/arrear", closeArrear)
	authorized.GET("/freeBooks", freeBooks)
	authorized.OPTIONS("/arrear", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
	authorized.OPTIONS("/freeBooks", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	r.POST("/auth", Login)
	r.OPTIONS("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
	r.GET("/auth", Refresh)

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
