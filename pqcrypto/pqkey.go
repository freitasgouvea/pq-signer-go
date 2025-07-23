package pqcrypto

import (
	"log"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

// GenerateKeypair generates a new PQC keypair using the Dilithium2 algorithm
func GenerateKeypair() ([]byte, []byte) {
	var sigAlg oqs.Signature

	err := sigAlg.Init("Dilithium2", nil)
	if err != nil {
		log.Fatalf("Error initializing PQC sig: %v", err)
	}

	pubKey, err := sigAlg.GenerateKeyPair()
	if err != nil {
		log.Fatalf("Error generating keypair: %v", err)
	}

	privKey := sigAlg.ExportSecretKey()
	if len(privKey) == 0 {
		log.Fatalf("Error exporting secret key: length zero")
	}

	return pubKey, privKey
}
