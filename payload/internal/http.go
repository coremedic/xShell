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

func (l *Link) NewRequest(data []byte) (*http.Request, error) {
	baseURL := l.Host + fmt.Sprintf("/result?id=%s", l.Id)
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", baseURL, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
