package transport

import (
	"encoding/binary"
	"io"

	"golang.org/x/crypto/ssh"
)

type Channel struct {
	ssh.Channel
}

/*
Write to ssh channel

Sends 8 byte header indicating total data size
*/
func (c *Channel) Write(data []byte) error {
	// Total data size
	dSize := uint64(len(data))
	// Build 8 byte header indicating total data size
	header := make([]byte, 8)
	// Put data size into header
	binary.BigEndian.AppendUint64(header[0:8], dSize)
	// Send 8 byte header
	_, err := c.Channel.Write(header)
	if err != nil {
		return err
	}
	// Send data
	_, err = c.Channel.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Channel) Read() ([]byte, error) {
	// Init buffer for 8 byte header
	header := make([]byte, 8)
	if _, err := io.ReadFull(c.Channel, header); err != nil {
		return nil, err
	}
	// Get total data size from header
	dSize := binary.BigEndian.Uint64(header[0:8])
	// Init buffer for read data
	rBuf := make([]byte, dSize)
	// Read data into buffer
	_, err := c.Channel.Read(rBuf)
	if err != nil {
		return nil, err
	}
	return rBuf, nil
}
