package internal

import (
	"bytes"
	"net/http"
	"net/url"
	"path"
)

type Link struct {
	Id   string
	Host string
	Key  []byte
}

var HttpLink Link
var client http.Client

func (l *Link) NewResultRequest(data []byte) (*http.Request, error) {
	u, err := url.Parse(l.Host)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "res")
	q := u.Query()
	q.Set("id", l.Id)
	u.RawQuery = q.Encode()

	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (l *Link) NewIdRequest() (*http.Request, error) {
	u, err := url.Parse(l.Host)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "id")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (l *Link) NewCmdRequest() (*http.Request, error) {
	u, err := url.Parse(l.Host)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "cmd")
	q := u.Query()
	q.Set("id", l.Id)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
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
