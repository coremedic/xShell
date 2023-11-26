package beacon

import (
	"bytes"
	"fmt"

	"github.com/enceve/crypto/serpent"
)

/*
Encrypt []byte data with Serpent block ciper and PKCS#7 padding

Return -> Encrypted []byte data, error
*/
//garble:controlflow flatten_passes=1 junk_jumps=1 block_splits=2
func SerpentEncrypt(data []byte, key []byte) ([]byte, error) {
	cipher, err := serpent.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Apply PKCS#7 padding
	data = pkcs7Pad(data, cipher.BlockSize())

	encDataBytes := make([]byte, len(data))
	for bs := 0; bs < len(data); bs += 16 {
		cipher.Encrypt(encDataBytes[bs:], data[bs:bs+16])
	}
	return encDataBytes, nil
}

/*
Decrypt Serpent block cipher encrypted data and remove PKCS#7 padding

Return -> Decrypted []byte data, error
*/
//garble:controlflow flatten_passes=1 junk_jumps=1 block_splits=2
func SerpentDecrypt(data []byte, key []byte) ([]byte, error) {
	cipher, err := serpent.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if remainder := len(data) % 16; remainder != 0 {
		return nil, fmt.Errorf("invalid block size")
	}

	// Decrypt each block
	dataBytes := make([]byte, len(data))
	for bs := 0; bs < len(data); bs += 16 {
		cipher.Decrypt(dataBytes[bs:], data[bs:bs+16])
	}

	// Remove PKCS#7 padding
	return pkcs7Unpad(dataBytes, cipher.BlockSize()), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

/*
Remove PKCS#7 padding
*/
func pkcs7Unpad(data []byte, blockSize int) []byte {
	length := len(data)
	unpadding := int(data[length-1])

	if unpadding > blockSize || unpadding == 0 {
		return data
	}

	padStart := length - unpadding
	for _, v := range data[padStart:] {
		if v != byte(unpadding) {
			return data
		}
	}

	return data[:padStart]
}
