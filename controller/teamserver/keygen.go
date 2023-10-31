package teamserver

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ed25519"
)

func GenHostKeys() (rsaPrivateKey []byte, rsaPublicKey []byte, ed25519PrivateKey []byte, ed25519PublicKey []byte, err error) {
	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rsaPrivateKeyDer := x509.MarshalPKCS1PrivateKey(rsaPrivKey)
	rsaPrivateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   rsaPrivateKeyDer,
	}
	rsaPrivateKey = pem.EncodeToMemory(&rsaPrivateKeyBlock)

	rsaPublicKeyDer, err := x509.MarshalPKIXPublicKey(&rsaPrivKey.PublicKey)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rsaPublicKeyBlock := pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   rsaPublicKeyDer,
	}
	rsaPublicKey = pem.EncodeToMemory(&rsaPublicKeyBlock)

	ed25519PublicKey, ed25519PrivKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(ed25519PrivKey)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	privateKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	ed25519PrivateKey = pem.EncodeToMemory(privateKeyBlock)

	return rsaPrivateKey, rsaPublicKey, ed25519PrivateKey, ed25519PublicKey, nil
}
