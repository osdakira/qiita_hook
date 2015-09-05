package chatsworth

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Chatsworth struct {
	RoomID   string
	APIToken string
}

const baseURL = "https://api.chatwork.com/v1/"

func (chatsworth Chatsworth) PostMessage(body string) {
	urlString := baseURL + "rooms/" + chatsworth.RoomID + "/messages"

	values := url.Values{}
	values.Add("body", body)

	req, _ := http.NewRequest("POST", urlString, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-ChatWorkToken", chatsworth.APIToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
}
