package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

var (
	key *ecdsa.PrivateKey
)

func SetupKey() {
	keyPath, exists := os.LookupEnv("JWT_PK_FILE")
	if !exists {
		log.Fatalln("[jwt/ecdsa] env var JWT_PK_FILE is not set.")
	} else {
		// getKey wil set the key variable
		getKey(keyPath)

		log.Println("[jwt/ecdsa] key loaded successfully")
	}
}

func loadKey(path string) {
	// Attempt to load the file
	dat, err := os.ReadFile(path)

	if err == nil {
		// Opened sucessfully
		block, _ := pem.Decode([]byte(dat))

		x509EncodedPrivKey := block.Bytes

		privateKey, err := x509.ParseECPrivateKey(x509EncodedPrivKey)

		if err != nil {
			log.Fatalln("[jwt/ecdsa] failed to parse keyfile")
		} else {
			key = privateKey
		}
	} else {
		log.Fatalln("[jwt/ecdsa] failed to open keyfile")
	}
}

func createKey(path string) {
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		log.Fatalln("[jwt/ecdsa] failed to create key")
	} else {
		bytes, marshalErr := x509.MarshalECPrivateKey(key)

		pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: bytes})

		if marshalErr != nil {
			log.Fatalln("[jwt/ecdsa] failed to marshal generated private key")
		}

		// Write the file, chmod 600 (owner r+w)
		writeErr := os.WriteFile(path, pemEncoded, 0600)

		if writeErr != nil {
			log.Fatalln("[jwt/ecdsa] failed to write generated key")
		}
	}
}

// Source: https://www.tutorialspoint.com/how-to-check-if-a-file-exists-in-golang
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return !info.IsDir()
	}
}

// Returns true if successful, panics if not.
func getKey(path string) {
	if !fileExists(path) {
		createKey(path)
	}
	loadKey(path)
}
