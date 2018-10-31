package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/reijo1337/online-library-rsoi2/clients"
	"github.com/reijo1337/online-library-rsoi2/mock_clients"
)

// TestUserArrearsNormal проверяет корректное выполнение запроса при нормальных параметрах запроса
func TestUserArrearsNormal(t *testing.T) {
	readerName := "Ivan Ivanov"
	page := 1
	size := 25

	resp := []ArrearResponse{
		ArrearResponse{
			ID:         1,
			ReaderID:   1,
			BookID:     1,
			BookName:   "Book name 1",
			BookAuthor: "Writer name",
			Start:      "start",
			End:        "end",
		},
		ArrearResponse{
			ID:         2,
			ReaderID:   1,
			BookID:     2,
			BookName:   "Book name 2",
			BookAuthor: "Writer name",
			Start:      "start",
			End:        "end",
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockBooksPartClient := mock_clients.NewMockBooksPartInterface(mockCtrl)
	MockReadersPartClient := mock_clients.NewMockReadersPartInterface(mockCtrl)
	MockArrearsPartClient := mock_clients.NewMockArrearsPartInterface(mockCtrl)

	BooksPartClient = MockBooksPartClient
	ArrearsPartClient = MockArrearsPartClient
	ReadersPartClient = MockReadersPartClient

	MockReadersPartClient.EXPECT().GetReaderByName(readerName).Return(clients.Reader{ID: 1, Name: readerName}, nil).Times(1)
	MockArrearsPartClient.EXPECT().GetArrearsPaging(int32(1), int32(page), int32(size)).Return(
		[]clients.Arrear{
			clients.Arrear{
				ID:       1,
				ReaderID: 1,
				BookID:   1,
				Start:    "start",
				End:      "end",
			},
			clients.Arrear{
				ID:       2,
				ReaderID: 1,
				BookID:   2,
				Start:    "start",
				End:      "end",
			},
		},
		nil,
	)
	MockBooksPartClient.EXPECT().GetBookByID(int32(1)).Return(
		&clients.Book{
			ID:   1,
			Name: "Book name 1",
			Free: false,
			Author: &clients.Writer{
				ID:   1,
				Name: "Writer name",
			},
		}, nil).Times(1)

	MockBooksPartClient.EXPECT().GetBookByID(int32(2)).Return(
		&clients.Book{
			ID:   2,
			Name: "Book name 2",
			Free: false,
			Author: &clients.Writer{
				ID:   1,
				Name: "Writer name",
			},
		}, nil).Times(1)

	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getUserArrears?name="+readerName+"&page="+strconv.Itoa(page)+"&size="+strconv.Itoa(size), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	var respReal []ArrearResponse
	err := json.Unmarshal([]byte(w.Body.String()), &respReal)
	assert.NoError(t, err)
	assert.ElementsMatch(t, respReal, resp)
}

//TestUserArrearsWrongQueryParams проверяет корректное выполнение запросов при некорректных параметрах
func TestUserArrearsWrongQueryParams(t *testing.T) {
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getUserArrears?name=Vladimir Putin&page=sas&size=25", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	req, _ = http.NewRequest("GET", "/getUserArrears?name=Vladimir Putin&page=1&size=sas", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	req, _ = http.NewRequest("GET", "/getUserArrears?name=Vladimir Putin&page=sas&size=sas", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

// TestUsersForBadResponsesFromServices проверяет корректное выполнение запроса если удаленные сервисы чего-то не могут
func TestUsersForBadResponsesFromServices(t *testing.T) {
	askedName := "GSPD"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockBooksPartClient := mock_clients.NewMockBooksPartInterface(mockCtrl)
	MockReadersPartClient := mock_clients.NewMockReadersPartInterface(mockCtrl)
	MockArrearsPartClient := mock_clients.NewMockArrearsPartInterface(mockCtrl)

	BooksPartClient = MockBooksPartClient
	ArrearsPartClient = MockArrearsPartClient
	ReadersPartClient = MockReadersPartClient

	MockReadersPartClient.EXPECT().GetReaderByName(askedName).Return(
		clients.Reader{},
		errors.New("Reader not exists"),
	).Times(1)
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getUserArrears?name="+askedName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
	assert.JSONEq(t, `{"error": "Reader not exists"}`, w.Body.String())

	MockReadersPartClient.EXPECT().GetReaderByName(askedName).Return(
		clients.Reader{ID: 1, Name: askedName}, nil,
	).Times(1)
	MockArrearsPartClient.EXPECT().GetArrearsPaging(int32(1), int32(1), int32(10)).Return(
		nil, errors.New("Can't get arrears"),
	).Times(1)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/getUserArrears?name="+askedName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.JSONEq(t, `{"error": "Can't get arrears"}`, w.Body.String())

	MockReadersPartClient.EXPECT().GetReaderByName(askedName).Return(
		clients.Reader{ID: 1, Name: askedName}, nil).Times(1)
	MockArrearsPartClient.EXPECT().GetArrearsPaging(int32(1), int32(1), int32(10)).Return(
		[]clients.Arrear{
			clients.Arrear{
				ID:       1,
				ReaderID: 1,
				BookID:   1,
				Start:    "start",
				End:      "end",
			},
			clients.Arrear{
				ID:       2,
				ReaderID: 1,
				BookID:   2,
				Start:    "start",
				End:      "end",
			},
		},
		nil,
	)
	MockBooksPartClient.EXPECT().GetBookByID(int32(1)).Return(
		nil, errors.New("There is no such book")).Times(1)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/getUserArrears?name="+askedName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
	assert.JSONEq(t, `{"error": "There is no such book"}`, w.Body.String())
}

// TestCorrectMakingNewArrear проверяет корректное выполнение запроса создания новой записи в читательском билете
func TestCorrectMakingNewArrear(t *testing.T) {
	var bookID int32
	bookID = 1
	name := "Viktor Kramov"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockBooksPartClient := mock_clients.NewMockBooksPartInterface(mockCtrl)
	MockReadersPartClient := mock_clients.NewMockReadersPartInterface(mockCtrl)
	MockArrearsPartClient := mock_clients.NewMockArrearsPartInterface(mockCtrl)

	BooksPartClient = MockBooksPartClient
	ArrearsPartClient = MockArrearsPartClient
	ReadersPartClient = MockReadersPartClient

	MockBooksPartClient.EXPECT().GetBookByID(bookID).Return(nil, nil).Times(1)
	MockReadersPartClient.EXPECT().GetReaderByName(name).Return(
		clients.Reader{ID: 1, Name: name}, nil,
	).Times(1)
	MockBooksPartClient.EXPECT().ChangeBookStatusByID(bookID, false).Return(nil).Times(1)
	MockArrearsPartClient.EXPECT().NewArrear(int32(1), bookID).Return(
		&clients.Arrear{
			ID:       1,
			ReaderID: 1,
			BookID:   bookID,
			Start:    "start",
			End:      "end",
		}, nil).Times(1)
	respExp := NewArrearResponse{
		ID:       1,
		ReaderID: 1,
		BookID:   bookID,
		Start:    "start",
		End:      "end",
	}

	postReq := clients.NewReaderWithArrearRequestBody{
		ReaderName: name,
		BookID:     bookID,
	}

	postReqJSON, err := json.Marshal(&postReq)
	assert.NoError(t, err)
	router := SetUpRouter()
	w := httptest.NewRecorder()
	fmt.Println(strings.NewReader(string(postReqJSON)))
	req, _ := http.NewRequest("POST", "/newArear", strings.NewReader(string(postReqJSON)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	var respReal NewArrearResponse
	err = json.Unmarshal([]byte(w.Body.String()), &respReal)
	assert.NoError(t, err)
	assert.Equal(t, respExp, respReal)
}
