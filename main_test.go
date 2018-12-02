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
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.JSONEq(t, `{"error": "Reader not exists"}`, w.Body.String())

	MockReadersPartClient.EXPECT().GetReaderByName(askedName).Return(
		clients.Reader{ID: 1, Name: askedName}, nil).AnyTimes()
	MockArrearsPartClient.EXPECT().GetArrearsPaging(int32(1), int32(1), int32(5)).Return(
		nil, errors.New("Can't get arrears"),
	).Times(1)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/getUserArrears?name="+askedName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.JSONEq(t, `{"error": "Can't get arrears"}`, w.Body.String())

	MockArrearsPartClient.EXPECT().GetArrearsPaging(int32(1), int32(1), int32(5)).Return(
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
		nil)
	MockBooksPartClient.EXPECT().GetBookByID(int32(1)).Return(
		nil, errors.New("There is no such book")).Times(1)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/getUserArrears?name="+askedName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
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

	MockBooksPartClient.EXPECT().GetBookByID(bookID).Return(&clients.Book{
		ID:     bookID,
		Name:   "Book Name",
		Author: &clients.Writer{Name: "Writer Name"},
	}, nil).Times(1)
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
		ID:         1,
		ReaderID:   1,
		BookID:     bookID,
		Start:      "start",
		End:        "end",
		BookName:   "Book Name",
		BookAuthor: "Writer Name",
	}

	postReq := clients.NewReaderWithArrearRequestBody{
		ReaderName: name,
		BookID:     bookID,
	}

	postReqJSON, err := json.Marshal(&postReq)
	assert.NoError(t, err)
	router := SetUpRouter()
	w := httptest.NewRecorder()
	fmt.Println("REQUEST", string(postReqJSON))
	req, _ := http.NewRequest("POST", "/arrear", strings.NewReader(string(postReqJSON)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	var respReal NewArrearResponse
	err = json.Unmarshal([]byte(w.Body.String()), &respReal)
	assert.NoError(t, err)
	assert.Equal(t, respExp, respReal)
}

// TestIncorrectMakingNewArrear проверяет корректное выполнение запроса создания новой записи при некорректных параметрах запроса
func TestIncorrectMakingNewArrear(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockBooksPartClient := mock_clients.NewMockBooksPartInterface(mockCtrl)
	MockReadersPartClient := mock_clients.NewMockReadersPartInterface(mockCtrl)
	MockArrearsPartClient := mock_clients.NewMockArrearsPartInterface(mockCtrl)

	BooksPartClient = MockBooksPartClient
	ArrearsPartClient = MockArrearsPartClient
	ReadersPartClient = MockReadersPartClient

	// Некорректное тело запроса
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/arrear", strings.NewReader("sas"))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)

	// Запрашиваемая книга отсутсвтует
	bookID := int32(1)
	MockBooksPartClient.EXPECT().GetBookByID(bookID).Return(
		nil, errors.New("There is no such book")).Times(1)
	postReq := clients.NewReaderWithArrearRequestBody{
		ReaderName: "doesn't matter",
		BookID:     bookID,
	}
	w = httptest.NewRecorder()
	postReqJSON, err := json.Marshal(&postReq)
	assert.NoError(t, err)
	req, _ = http.NewRequest("POST", "/arrear", strings.NewReader(string(postReqJSON)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.JSONEq(t, `{"error": "There is no such book"}`, w.Body.String())

	// Запись на несуществующего читателя
	readerName := "Test Reader"
	MockBooksPartClient.EXPECT().GetBookByID(bookID).Return(nil, nil).AnyTimes()
	MockReadersPartClient.EXPECT().GetReaderByName(readerName).Return(
		clients.Reader{}, errors.New("There is no such reader")).Times(1)
	postReq = clients.NewReaderWithArrearRequestBody{
		ReaderName: readerName,
		BookID:     bookID,
	}
	w = httptest.NewRecorder()
	postReqJSON, err = json.Marshal(&postReq)
	assert.NoError(t, err)
	req, _ = http.NewRequest("POST", "/arrear", strings.NewReader(string(postReqJSON)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.JSONEq(t, `{"error": "There is no such reader"}`, w.Body.String())

	// Книга уже занята
	readerID := int32(1)
	MockReadersPartClient.EXPECT().GetReaderByName(readerName).Return(
		clients.Reader{ID: readerID, Name: readerName}, nil).AnyTimes()
	MockBooksPartClient.EXPECT().ChangeBookStatusByID(bookID, false).Return(
		errors.New("Can't change book status")).Times(1)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/arrear", strings.NewReader(string(postReqJSON)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
	assert.JSONEq(t, `{"error": "Can't change book status"}`, w.Body.String())

	// Проблемы при регистрации записи
	MockBooksPartClient.EXPECT().ChangeBookStatusByID(bookID, false).Return(nil).Times(1)
	MockArrearsPartClient.EXPECT().NewArrear(readerID, bookID).Return(
		nil, errors.New("Can't register new arrear")).Times(1)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/arrear", strings.NewReader(string(postReqJSON)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
	assert.JSONEq(t, `{"error": "Can't register new arrear"}`, w.Body.String())
}

// TestCorrectCloseArrear проверяет корректное выполнение запроса на возврат книги при корректных параметрах запроса
func TestCorrectCloseArrear(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockBooksPartClient := mock_clients.NewMockBooksPartInterface(mockCtrl)
	MockReadersPartClient := mock_clients.NewMockReadersPartInterface(mockCtrl)
	MockArrearsPartClient := mock_clients.NewMockArrearsPartInterface(mockCtrl)

	BooksPartClient = MockBooksPartClient
	ArrearsPartClient = MockArrearsPartClient
	ReadersPartClient = MockReadersPartClient

	arrearID := int32(1)
	readerID := int32(1)
	bookID := int32(1)
	arrearIDstring := fmt.Sprint(arrearID)
	MockArrearsPartClient.EXPECT().GetArrearByID(arrearID).Return(
		&clients.Arrear{
			ID:       arrearID,
			ReaderID: readerID,
			BookID:   bookID,
			Start:    "start",
			End:      "end",
		}, nil).Times(1)
	MockArrearsPartClient.EXPECT().CloseArrearByID(arrearID).Return(nil).Times(1)
	MockBooksPartClient.EXPECT().ChangeBookStatusByID(bookID, true).Return(nil).Times(1)

	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/arrear?id="+arrearIDstring, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.JSONEq(t, `{"ok": "ok"}`, w.Body.String())
}

// TestIncorrectMakingNewArrear проверяет корректное выполнение запроса на возврат книги при некорректных параметрах запроса
func TestIncorrectCloseArrear(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockBooksPartClient := mock_clients.NewMockBooksPartInterface(mockCtrl)
	MockReadersPartClient := mock_clients.NewMockReadersPartInterface(mockCtrl)
	MockArrearsPartClient := mock_clients.NewMockArrearsPartInterface(mockCtrl)

	BooksPartClient = MockBooksPartClient
	ArrearsPartClient = MockArrearsPartClient
	ReadersPartClient = MockReadersPartClient

	arrearID := int32(1)
	readerID := int32(1)
	bookID := int32(1)
	arrearIDstring := fmt.Sprint(arrearID)
	incorArrearIDstring := "sas"

	// Некорректный ID записи
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/arrear?id="+incorArrearIDstring, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)

	// Удаляемая запись не существует
	MockArrearsPartClient.EXPECT().GetArrearByID(arrearID).Return(nil, errors.New("There is no such arrear")).Times(1)
	router = SetUpRouter()
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/arrear?id="+arrearIDstring, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.JSONEq(t, `{"error": "There is no such arrear"}`, w.Body.String())

	// Проблемы при закрытии записи
	MockArrearsPartClient.EXPECT().GetArrearByID(arrearID).Return(
		&clients.Arrear{
			ID:       arrearID,
			ReaderID: readerID,
			BookID:   bookID,
			Start:    "start",
			End:      "end",
		}, nil).AnyTimes()
	MockArrearsPartClient.EXPECT().CloseArrearByID(arrearID).Return(
		errors.New("Some error while closing arrear"))
	router = SetUpRouter()
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/arrear?id="+arrearIDstring, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
	assert.JSONEq(t, `{"error": "Some error while closing arrear"}`, w.Body.String())

	// Проблемы при возврата книге статуса FREE
	MockArrearsPartClient.EXPECT().CloseArrearByID(arrearID).Return(nil).AnyTimes()
	MockBooksPartClient.EXPECT().ChangeBookStatusByID(bookID, true).Return(errors.New("Can't change book status"))
	router = SetUpRouter()
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/arrear?id="+arrearIDstring, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
	assert.JSONEq(t, `{"error": "Can't change book status"}`, w.Body.String())
}
