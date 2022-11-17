package encryption

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	key := "album"
	value := "bottle it in"
	secretKey := "adjfhoaudhvc90234hrhjb2i34ob"
	// Encrypt the key and value
	encryptedKey, err := Encrypt(key, secretKey)
	if err != nil {
		t.Fatalf("unexpected error while encrypting key: %q", err)
	}
	print(encryptedKey)
	encryptedVal, err := Encrypt(value, secretKey)
	if err != nil {
		t.Fatalf("unexpected error while encrypting value: %q", err)
	}
	// Decrypt the key and value
	decryptedKey, err := Decrypt(encryptedKey, secretKey)
	if err != nil {
		t.Fatalf("unexpected error while decrypting key: %q", err)
	}
	decryptedVal, err := Decrypt(encryptedVal, secretKey)
	if err != nil {
		t.Fatalf("unexpected error while decrypting value: %q", err)
	}
	// Check the values are the same as expected
	if key != string(decryptedKey) {
		t.Fatalf("error occured during encryption process, expected %s, instead got %s",
			key, string(decryptedKey),
		)
	}
	if value != string(decryptedVal) {
		t.Fatalf("error occured during encryption process, expected %s, instead got %s",
			value, string(decryptedVal),
		)
	}
}
