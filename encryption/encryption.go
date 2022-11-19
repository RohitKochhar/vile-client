package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"io"
	"strings"
)

// hashInput hashes the input bytes using the md5 Sum method
func getGCMInstance(secretKey string) (cipher.AEAD, error) {
	hashedSecretKey := sha256.Sum256([]byte(secretKey))
	aesBlock, err := aes.NewCipher([]byte(hashedSecretKey[:]))
	if err != nil {
		return nil, fmt.Errorf("unexpected error while hasing secret key: %q", err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while creating GCM instance: %q", err)
	}
	return gcmInstance, nil
}

// Encrypt encrypts the value provided using the provided secretKey
func Encrypt(value string, secretKey string) (string, error) {
	gcmInstance, err := getGCMInstance(secretKey)
	if err != nil {
		return "", fmt.Errorf("unexpected error while creating GCM instance: %q", err)
	}
	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(strings.NewReader(""), nonce)
	cipheredText := gcmInstance.Seal(nonce, nonce, []byte(value), nil)
	return base32.HexEncoding.EncodeToString(cipheredText), nil
}

// Decrypt decrypts the value provided using the provided secretKey
func Decrypt(cipheredInput string, secretKey string) ([]byte, error) {
	ciphered, err := base32.HexEncoding.DecodeString(cipheredInput)
	if err != nil {
		return nil, fmt.Errorf("error while decoding ciphered text: %q", err)
	}
	gcmInstance, err := getGCMInstance(secretKey)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while creating GCM instance: %q", err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]
	decrypted, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while decrypting value: %q", err)
	}
	return decrypted, nil
}
