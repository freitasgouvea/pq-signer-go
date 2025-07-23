package pqcrypto

import (
	"errors"
	"os"
)

// SaveKeypairToFile saves the public and private keys to specified file paths.
func SaveKeypairToFile(pubKey, privKey []byte, pubPath, privPath string) error {
	if err := os.WriteFile(pubPath, pubKey, 0600); err != nil {
		return err
	}
	if err := os.WriteFile(privPath, privKey, 0600); err != nil {
		return err
	}
	return nil
}

// LoadKeypairFromFile loads the public and private keys from specified file paths.
func LoadKeypairFromFile(pubPath, privPath string) ([]byte, []byte, error) {
	pubKey, err := os.ReadFile(pubPath)
	if err != nil {
		return nil, nil, err
	}
	privKey, err := os.ReadFile(privPath)
	if err != nil {
		return nil, nil, err
	}
	if len(pubKey) == 0 || len(privKey) == 0 {
		return nil, nil, errors.New("key file is empty")
	}
	return pubKey, privKey, nil
}
