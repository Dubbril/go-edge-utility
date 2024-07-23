package utils

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func EncryptWithBase64Key(plainText, key string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Pad the plaintext if necessary
	padding := aes.BlockSize - len(plainTextBytes)%aes.BlockSize
	if padding != 0 {
		for i := 0; i < padding; i++ {
			plainTextBytes = append(plainTextBytes, byte(padding))
		}
	}

	cipherText := make([]byte, len(plainTextBytes))
	// Encrypt each block separately
	for bs, be := 0, block.BlockSize(); bs < len(plainTextBytes); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(cipherText[bs:be], plainTextBytes[bs:be])
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptWithBase64Key(encrypted, key string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Encrypt each block separately
	decrypted := make([]byte, len(encryptedBytes))
	for bs, be := 0, block.BlockSize(); bs < len(encryptedBytes); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(decrypted[bs:be], encryptedBytes[bs:be])
	}

	plaintext, err := removePKCS7Padding(decrypted)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func EncryptWithHexKey(plainText, key string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Pad the plaintext if necessary
	padding := aes.BlockSize - len(plainTextBytes)%aes.BlockSize
	if padding != 0 {
		for i := 0; i < padding; i++ {
			plainTextBytes = append(plainTextBytes, byte(padding))
		}
	}

	cipherText := make([]byte, len(plainTextBytes))
	// Encrypt each block separately
	for bs, be := 0, block.BlockSize(); bs < len(plainTextBytes); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(cipherText[bs:be], plainTextBytes[bs:be])
	}

	return hex.EncodeToString(cipherText), nil
}

func DecryptWithHexKey(encrypted, key string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	encryptedBytes, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Encrypt each block separately
	decrypted := make([]byte, len(encryptedBytes))
	for bs, be := 0, block.BlockSize(); bs < len(encryptedBytes); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(decrypted[bs:be], encryptedBytes[bs:be])
	}

	plaintext, err := removePKCS7Padding(decrypted)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func removePKCS7Padding(data []byte) ([]byte, error) {
	padLen := int(data[len(data)-1])
	if padLen > len(data) || padLen > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding length")
	}
	return data[:len(data)-padLen], nil
}
