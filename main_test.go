package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/reijo1337/online-library-rsoi2/clients"
	"github.com/reijo1337/online-library-rsoi2/mock_clients"
)

// TestUserArrearsNormal проверяет корректное выполнение запроса при нормальных параметрах запроса
func TestUserArrearsNormal(t *testing.T) {
	readerName := "Ivan Ivanov"
	page := 1
	size := 25

	mockCtrl := gomock.NewController(t)

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
	if w.Code != http.StatusOK {
		t.Log("Code:", w.Code)
		t.Log("Body:", w.Body.String())
		t.Fail()
	}

}
