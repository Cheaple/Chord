package chord

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"path"
	"time"
)


func (n *Node) startTlsConfig() error {
	// Check certificate existence
	cerPath := n.getCerticatePath()
	keyPath := n.getPrivateKeyPath()
	if _, err := os.Stat(cerPath); err == nil {
		if _, err := os.Stat(keyPath); err == nil {
			return nil
		}
	}

	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}

	// Build certificate template
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: string(n.Address)},  // certificate name: node's address
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),  // period of validity: 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Create certificate using the template
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Error creating certificate: %v", err)
	}

	// Write certificate to local file
	certFile, err := os.Create(cerPath)
	if err != nil {
		log.Fatalf("Error creating cert file: %v", err)
	}
	defer certFile.Close()
	pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	// Write private key to local file
	keyFile, err := os.Create(keyPath)
	if err != nil {
		log.Fatalf("Error creating key file: %v", err)
	}
	defer keyFile.Close()
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Error marshaling private key: %v", err)
	}
	pem.Encode(keyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})

	fmt.Println("Certificates generated successfully.")
	return nil
}

func (n *Node) getCerticatePath() string {
	return path.Join(n.baseDir, "server.crt")
}

func (n *Node) getPrivateKeyPath() string {
	return path.Join(n.baseDir, "server.key")
}