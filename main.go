package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/freitasgouvea/pq-signer/pqcrypto"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "pq-signer",
		Short: "Post-Quantum Signer",
	}

	var genKeyCmd = &cobra.Command{
		Use:   "genkey",
		Short: "Generate a PQC keypair and save to files",
		Run: func(cmd *cobra.Command, args []string) {
			pubKey, privKey := pqcrypto.GenerateKeypair()

			// Ensure keys directory exists
			err := os.MkdirAll("keys", 0700)
			if err != nil {
				log.Fatalf("Failed to create keys directory: %v", err)
			}

			err = pqcrypto.SaveKeypairToFile(pubKey, privKey, "keys/pub.key", "keys/priv.key")
			if err != nil {
				log.Fatalf("Failed to save keys: %v", err)
			}
			fmt.Println("✅ Keys generated and saved to keys/pub.key and keys/priv.key")
		},
	}

	var signCmd = &cobra.Command{
		Use:   "sign [message]",
		Short: "Sign a message using the private key",
		Args:  cobra.ExactArgs(1), // expect exactly 1 argument: message to sign
		Run: func(cmd *cobra.Command, args []string) {
			pubKey, privKey, err := pqcrypto.LoadKeypairFromFile("keys/pub.key", "keys/priv.key")
			if err != nil {
				log.Fatalf("Failed to load keys: %v", err)
			}

			message := []byte(args[0])

			signature := pqcrypto.SignMessage(privKey, message)
			fmt.Printf("Signature (hex):\n%x\n", signature)

			if pqcrypto.VerifySignature(pubKey, message, signature) {
				fmt.Println("✅ Signature verified!")
			} else {
				fmt.Println("❌ Signature verification failed.")
			}
		},
	}

	var verifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "Verify a signature by passing message and signature as arguments",
		Args:  cobra.ExactArgs(2), // expect exactly 2 arguments: message and signature (hex)
		Run: func(cmd *cobra.Command, args []string) {
			pubKey, _, err := pqcrypto.LoadKeypairFromFile("keys/pub.key", "keys/priv.key")
			if err != nil {
				log.Fatalf("Failed to load public key: %v", err)
			}

			message := []byte(args[0])

			sigHex := args[1]
			signature, err := hex.DecodeString(sigHex)
			if err != nil {
				log.Fatalf("Failed to decode signature hex: %v", err)
			}

			if pqcrypto.VerifySignature(pubKey, message, signature) {
				fmt.Println("✅ Signature verified!")
			} else {
				fmt.Println("❌ Signature verification failed.")
			}
		},
	}

	rootCmd.AddCommand(genKeyCmd)
	rootCmd.AddCommand(signCmd)
	rootCmd.AddCommand(verifyCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
