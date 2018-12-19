package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AuthPartInterface interface {
	GetToken(user *User) (*Tokens, error)
	RefreshToken(refreshToken string) (*Tokens, error)
}

type AuthPart struct {
	url string
}

func NewAuthPart() (*AuthPart, error) {
	log.Println("Auth Client: Setting new authorization client")
	authURL := os.Getenv("AUTH")
	authURL = "http://127.0.0.1:8084"
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
	req, err := http.NewRequest("POST", ap.url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Auth Client: Can't send request: ", err.Error())
		return nil, err
	}
	if err != nil {
		log.Println("Auth Client: Can't send request: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if resp.StatusCode != http.StatusOK {
		errMsg := &ErrorMessage{}
		json.Unmarshal(body, errMsg)
		return nil, errors.New(errMsg.Error)
	}
	tokens := &Tokens{}
	if err = json.Unmarshal(body, tokens); err != nil {
		log.Println("Auth Client: Can't Unmarshal response: ", err.Error())
		return nil, err
	}
	return tokens, nil
}

func (ap *AuthPart) RefreshToken(RefreshToken string) (*Tokens, error) {
	log.Println("Auth Client: Request for refresh token")
	resp, err := http.Get(ap.url + "?refresh_token=" + RefreshToken)
	if err != nil {
		log.Println("Auth Client: Can't reshresh token: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	tokens := &Tokens{}
	if err = json.Unmarshal(body, tokens); err != nil {
		log.Println("Auth Client: Can't Unmarshal response: ", err.Error())
		return nil, err
	}
	return tokens, nil
}
