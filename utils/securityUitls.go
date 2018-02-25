package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var (
	ErrCipherTextTooShort = errors.New("Ciphertext block is too short.")
)

func Encrypt(key, content string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(content))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(cipherText[aes.BlockSize:], []byte(content))
	result := base64.URLEncoding.EncodeToString(cipherText)
	return result, nil
}

func Decrypt(key, encrypt string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", ErrCipherTextTooShort
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
