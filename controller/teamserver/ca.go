package teamserver

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

/*
Generate Certificate Authority (CA) certificate and key

Will generate a certificate and key for use as a CA.
Cert/Key will be used to sign client certificates for mTLS authentication.

Return -> PEM encoded cert, PEM endcoded key, error
*/
func GenCACert() ([]byte, []byte, error) {
	// Generate ecdsa key
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Cert is valid for 1 year
	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	// x509 cert template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Big Ballz CA"}, // Nuff said...
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Create our x509 cert
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// Get ecdsa private key from cert
	privDER, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}

	// Return PEM encoded cert and key
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privDER})
	return certPEM, keyPEM, nil
}

/*
Generate Client certificate and key from CA cert/key

Will generate a client cert/key for use with client mTLS connections.

Args -> CA Certificate, CA Key, Username

Return -> PEM encoded cert and key, error
*/
func GenClientCert(caCert *x509.Certificate, caKey *ecdsa.PrivateKey, username string) ([]byte, error) {
	// Generate ecdsa key
	clientPriv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	// Cert is valid for 1 year
	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	// x509 cert template
	clientCertTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"xShell Client"},
			CommonName:   username,
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// Create our x509 cert
	clientCertDER, err := x509.CreateCertificate(rand.Reader, &clientCertTemplate, caCert, &clientPriv.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	// Get ecdsa private key from cert
	clientPrivDER, err := x509.MarshalECPrivateKey(clientPriv)
	if err != nil {
		return nil, err
	}

	// Return PEM encoded cert and key
	clientCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: clientCertDER})
	clientKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: clientPrivDER})
	combinedPEM := append(clientCertPEM, clientKeyPEM...)
	return combinedPEM, nil
}

/*
Generate TeamServer certificate and key from CA cert/key

Will generate a TeamServer cert/key for use with client mTLS connections.

Args -> CA Certificate, CA Key

Return -> PEM encoded cert, PEM endcoded key, error
*/
func GenTeamServerCert(caCert *x509.Certificate, caKey *ecdsa.PrivateKey) ([]byte, []byte, error) {
	// Generate ecdsa key
	clientPriv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Cert is valid for 1 year
	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	// x509 cert template
	clientCertTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"xShell TeamServer"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// Create our x509 cert
	clientCertDER, err := x509.CreateCertificate(rand.Reader, &clientCertTemplate, caCert, &clientPriv.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	// Get ecdsa private key from cert
	clientPrivDER, err := x509.MarshalECPrivateKey(clientPriv)
	if err != nil {
		return nil, nil, err
	}

	// Return PEM encoded cert and key
	clientCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: clientCertDER})
	clientKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: clientPrivDER})
	return clientCertPEM, clientKeyPEM, nil
}
