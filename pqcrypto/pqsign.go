package pqcrypto

import (
	"log"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

// SignMessage signs a message using the provided private key
func SignMessage(priv []byte, message []byte) []byte {
	var sigAlg oqs.Signature

	// Initialize with algorithm and private key
	if err := sigAlg.Init("Dilithium2", priv); err != nil {
		log.Fatalf("Failed to initialize signature algorithm with private key: %v", err)
	}
	defer sigAlg.Clean()

	// Sign the message using the private key now loaded in sigAlg
	signature, err := sigAlg.Sign(message)
	if err != nil {
		log.Fatalf("Failed to sign message: %v", err)
	}

	return signature
}

// VerifySignature verifies a signature using the public key and message
func VerifySignature(pub []byte, message []byte, signature []byte) bool {
	var sigAlg oqs.Signature

	err := sigAlg.Init("Dilithium2", nil)
	if err != nil {
		log.Fatalf("Failed to initialize signature algorithm: %v", err)
	}
	defer sigAlg.Clean()

	ok, err := sigAlg.Verify(message, signature, pub)
	if err != nil {
		log.Fatalf("Failed to verify signature: %v", err)
	}

	return ok
}
