package internal

import (
	"bytes"
	"fmt"
	"net/http"
)

type Link struct {
	Id   string
	Host string
	Key  []byte
}

var HttpLink Link
var client http.Client

func (l *Link) NewResultRequest(data []byte) (*http.Request, error) {
	baseURL := l.Host + fmt.Sprintf("/res?id=%s", l.Id)
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", baseURL, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (l *Link) NewIdRequest() (*http.Request, error) {
	baseURL := l.Host + "/id"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (l *Link) NewCmdRequest() (*http.Request, error) {
	baseURL := l.Host + fmt.Sprintf("/cmd?id=%s", l.Id)
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func SendRequest(req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
