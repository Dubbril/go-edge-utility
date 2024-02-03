package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

const (
	ALGORITHM = "AES"
	CIPHER    = "AES"
)

func EncryptWithHexKey(text, key string) (string, error) {
	decodedKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", err
	}

	// Padding the plaintext to a multiple of block size
	text = addPKCS7Padding(text, block.BlockSize())

	ciphertext := make([]byte, len(text))
	blockMode := cipher.NewCBCEncrypter(block, decodedKey[:block.BlockSize()])
	blockMode.CryptBlocks(ciphertext, []byte(text))

	return hex.EncodeToString(ciphertext), nil
}

func EncryptWithBase64Key(text, key string) (string, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", err
	}

	// Padding the plaintext to a multiple of block size
	text = addPKCS7Padding(text, block.BlockSize())

	ciphertext := make([]byte, len(text))
	blockMode := cipher.NewCBCEncrypter(block, decodedKey[:block.BlockSize()])
	blockMode.CryptBlocks(ciphertext, []byte(text))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptWithHexKey(encrypted, key string) (string, error) {
	decodedKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	decodedEncrypted, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", err
	}

	if len(decodedEncrypted)%block.BlockSize() != 0 {
		return "", errors.New("encrypted text is not a multiple of the block size")
	}

	blockMode := cipher.NewCBCDecrypter(block, decodedKey[:block.BlockSize()])
	blockMode.CryptBlocks(decodedEncrypted, decodedEncrypted)

	// Remove PKCS7 padding
	decryptedText := removePKCS7Padding(decodedEncrypted)

	return string(decryptedText), nil
}

func DecryptWithBase64Key(encrypted, key string) (string, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	decodedEncrypted, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", err
	}

	if len(decodedEncrypted)%block.BlockSize() != 0 {
		return "", errors.New("encrypted text is not a multiple of the block size")
	}

	blockMode := cipher.NewCBCDecrypter(block, decodedKey[:block.BlockSize()])
	blockMode.CryptBlocks(decodedEncrypted, decodedEncrypted)

	// Remove PKCS7 padding
	decryptedText := removePKCS7Padding(decodedEncrypted)

	return string(decryptedText), nil
}

func addPKCS7Padding(text string, blockSize int) string {
	padding := blockSize - (len(text) % blockSize)
	paddingText := strings.Repeat(string(byte(padding)), padding)
	return text + paddingText
}

func removePKCS7Padding(text []byte) []byte {
	padding := int(text[len(text)-1])
	return text[:len(text)-padding]
}
