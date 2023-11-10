package teamserver

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
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

	// Generate a unique serial number for the certificate
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, err
	}

	// x509 cert template
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Big Ballz CA"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage: x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		IsCA:     true,
		// ExtKeyUsage can be omitted for a CA certificate
		// ExtKeyUsage:           []x509.ExtKeyUsage{},
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

Return -> PEM encoded fullchain certificate (CA cert, Client cert, Client key), error
*/
func GenClientCert(caCert *x509.Certificate, caKey *ecdsa.PrivateKey, username string) ([]byte, error) {
	// Generate an ECDSA key pair for the client
	clientPriv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	// Certificate is valid for 1 year
	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	// Generate a unique serial number for the certificate
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}

	// x509 certificate template for the client
	clientCertTemplate := x509.Certificate{
		SerialNumber: serialNumber,
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

	// Create the client certificate signed by the CA
	clientCertDER, err := x509.CreateCertificate(rand.Reader, &clientCertTemplate, caCert, &clientPriv.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	// Marshal the client private key into DER format
	clientPrivDER, err := x509.MarshalECPrivateKey(clientPriv)
	if err != nil {
		return nil, err
	}

	// Encode the client certificate to PEM
	clientCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: clientCertDER})
	// Encode the client private key to PEM
	clientKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: clientPrivDER})
	// Optionally, encode the CA certificate to PEM if it should be included
	caCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caCert.Raw})

	// Concatenate the client cert, client key, and optionally the CA cert into one PEM file
	fullChainPEM := append(clientCertPEM, clientKeyPEM...)
	fullChainPEM = append(fullChainPEM, caCertPEM...) // Append CA cert if needed

	return fullChainPEM, nil
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

	// Fetch host IP addresses
	ips, err := getHostIPAddresses()
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
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IPAddresses:           ips,
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

/*
Get ip addresses of all interfaces on host
*/
func getHostIPAddresses() ([]net.IP, error) {
	var ips []net.IP

	// Include localhost
	ips = append(ips, net.ParseIP("127.0.0.1"))

	// Get all network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// Iterate over all network interfaces
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		// Iterate over all addresses for interface
		for _, addr := range addrs {
			var ip net.IP

			// Check if it's an IPNet type (not a loopback)
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Check if the IP is non-loopback and non-link-local
			if ip != nil && !ip.IsLoopback() && !ip.IsLinkLocalUnicast() {
				ips = append(ips, ip)
			}
		}
	}

	return ips, nil
}