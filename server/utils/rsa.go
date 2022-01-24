package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

var privateKey, _ = rsa.GenerateKey(rand.Reader, 1024)

func Encryption(password string) ([]byte, error) {

	publicKey := privateKey.PublicKey

	encryptionBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte(password),
		nil)

	if err != nil {
		return nil, err
	}

	return encryptionBytes, nil
}

func Decryption(password string) ([]byte, error) {

	encryptedBytes := []byte(password)

	decryptedBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		encryptedBytes,
		nil)

	if err != nil {
		return nil, err
	}

	// decryptedStr := string(decryptedBytes)

	return decryptedBytes, nil
}

func Compare(encrypted, password []byte) bool {

	decryptedBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		encrypted,
		nil)

	if err != nil {
		return false
	}

	res := bytes.Compare(decryptedBytes, password)

	if res == 0 {
		return true
	} else {
		return false
	}
}
