package internal

import (
	"bytes"
	"fmt"

	"github.com/enceve/crypto/serpent"
)

func SerpentEncrypt(data []byte, key []byte) ([]byte, error) {
	cipher, err := serpent.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// pad data if necessary
	if remainder := len(data) % 16; remainder != 0 {
		padding := make([]byte, 16-remainder)
		data = append(data, padding...)
	}
	encDataBytes := make([]byte, len(data))
	for bs := 0; bs < len(data); bs += 16 {
		cipher.Encrypt(encDataBytes[bs:], data[bs:bs+16])
	}
	return encDataBytes, nil
}

func SerpentDecrypt(data []byte, key []byte) ([]byte, error) {
	if data == nil || bytes.Equal(data, []byte("")) {
		return nil, fmt.Errorf("null data")
	}
	cipher, err := serpent.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if remainder := len(data) % 16; remainder != 0 {
		return nil, fmt.Errorf("failed to decrypt data invalid block size")
	}

	// decrypt each block
	dataBytes := make([]byte, len(data))
	for bs := 0; bs < len(data); bs += 16 {
		cipher.Decrypt(dataBytes[bs:], data[bs:bs+16])
	}
	dataBytes = removePadding(dataBytes)
	return dataBytes, nil
}

func removePadding(data []byte) []byte {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != 0x00 {
			return data[:i+1]
		}
	}
	return nil
}
