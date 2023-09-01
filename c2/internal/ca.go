package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

type X509Cert struct {
	Orgs     []string
	Names    []string
	Ips      []string
	CertFile string
	KeyFile  string
}

// Generate self-signed certificate and key files
func GenerateCertificate(cert *X509Cert) error {
	// Generate RSA private key
	privKey, err := rsa.GenerateKey(rand.Reader, 2048) // TODO optional 4096 bit keys
	if err != nil {
		return err
	}
	ipAddrs := func(ips []string) []net.IP {
		pIps := make([]net.IP, len(ips))
		for i, npIp := range ips {
			pIps[i] = net.ParseIP(npIp)
		}
		return pIps
	}
	// Generate X.509 certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: cert.Orgs,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:              cert.Names,
		IPAddresses:           ipAddrs(cert.Ips),
	}
	// Generate X.509 certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		return err
	}
	// Encode X.509 certificate in PEM format
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	// Encode private key in PEM format
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})
	// Write certificate and private key files to disk
	err = os.WriteFile(cert.CertFile, certPEM, 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(cert.KeyFile, keyPEM, 0600)
	if err != nil {
		os.Remove(cert.CertFile)
		return err
	}
	return nil
}

func ValidateCertificate(cert *X509Cert) error {
	keyPair, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		if os.IsNotExist(err) {
			err := GenerateCertificate(cert)
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	} else {
		x509Cert, err := x509.ParseCertificate(keyPair.Certificate[0])
		if err != nil {
			return err
		}
		if x509Cert.NotAfter.Before(time.Now()) {
			err := GenerateCertificate(cert)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
