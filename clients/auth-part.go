package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AuthPartInterface interface {
	GetToken(user *User) (*Tokens, error)
}

type AuthPart struct {
	url string
}

func NewAuthPart() (*AuthPart, error) {
	log.Println("Auth Client: Setting new authorization client")
	authURL := os.Getenv("AUTH")
	return &AuthPart{
		url: authURL,
	}, nil
}

func (ap *AuthPart) GetToken(user *User) (*Tokens, error) {
	log.Println("Auth Client: Request for new token for ", user.Login)
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Println("Auth Client: Can't marshal request: ", err.Error())
		return nil, err
	}
	log.Println(string(jsonStr))
	// req, err := http.NewRequest("POST", ap.url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println("Auth Client: Can't send request: ", err.Error())
	// 	return nil, err
	// }
	resp, err := http.Post(ap.url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("Auth Client: Can't send request: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	var tokens *Tokens
	if err = json.Unmarshal(body, tokens); err != nil {
		log.Println("Auth Client: Can't Unmarshal response: ", err.Error())
		return nil, err
	}
	return tokens, nil
}
