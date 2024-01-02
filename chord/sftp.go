package chord

// For use file transfer using SFTP

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/pkg/sftp"
	"log"
	"os"
	"path"
)

func (n *Node) startSFTPService() error {
	n.initSSHConfig()

	// Load the private key file
	privateKeyBytes, err := ioutil.ReadFile(n.getPrivateKeyPath())
	if err != nil {
		log.Fatalf("Error loading private key: %v", err)
	}

	// Parse the private key
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	// SSH server configuration
	sshConfig := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			// Implement user authentication logic here, this example leaves it empty
			return nil, nil
		},
		// Allow public key authentication
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			// This example allows public key authentication, it's left empty for customization
			return nil, nil
		},
	}

	// Add the private key to the server configuration
	sshConfig.AddHostKey(privateKey)

	// Start SSH server
	listener, err := ssh.Listen("tcp", "localhost:2222", sshConfig)
	if err != nil {
		log.Fatalf("Error starting SSH server: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting incoming connection: %v", err)
		}

		// Create an SFTP server
		sftpServer, err := sftp.NewServer(
			conn,
		)
		if err != nil {
			log.Fatalf("Error creating SFTP server: %v", err)
		}

		// Start the SFTP server
		go func() {
			if err := sftpServer.Serve(); err != nil {
				log.Printf("SFTP server exited with error: %v", err)
			}
		}()
	}

	return nil
}

//
// Initialize SSH key pairs
//
func (n *Node) initSSHConfig() error {
	// Check SSH key existence
	publicKeyPath := n.getPublicKeyPath()
	privateKeyPath := n.getPrivateKeyPath()
	if _, err := os.Stat(publicKeyPath); err == nil {
		if _, err := os.Stat(privateKeyPath); err == nil {
			return nil
		}
	}

	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}

	// Serialize private key to ASN.1 DER encoding
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	// Create private key PEM block
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	// Write private key PEM block to a file
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		log.Fatalf("Error creating private key file: %v", err)
	}
	defer privateKeyFile.Close()

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		log.Fatalf("Error writing private key to file: %v", err)
	}

	// Generate public key
	publicKey := privateKey.PublicKey

	// Serialize public key to ASN.1 DER encoding
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatalf("Error marshaling public key: %v", err)
	}

	// Create public key PEM block
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// Write public key PEM block to a file
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		log.Fatalf("Error creating public key file: %v", err)
	}
	defer publicKeyFile.Close()

	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		log.Fatalf("Error writing public key to file: %v", err)
	}

	return nil
}

func (n *Node) getPublicKeyPath() string {
	return path.Join(n.baseDir, fmt.Sprintf("public-%d.pem", n.Id))
}

func (n *Node) getPrivateKeyPath() string {
	return path.Join(n.baseDir, fmt.Sprintf("private-%d.pem", n.Id))
}