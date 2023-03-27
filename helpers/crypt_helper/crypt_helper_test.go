package crypt_helper

import "testing"

func TestEncryption(t *testing.T) {
	key := []byte("mysecretkey12345")
	plaintext := []byte("hello world")

	// Encrypt the plaintext
	ciphertext, err := Encrypt(key, plaintext)
	if err != nil {
		t.Errorf("Error encrypting: %s", err)
	}

	// Decrypt the ciphertext
	decrypted, err := Decrypt(key, ciphertext)
	if err != nil {
		t.Errorf("Error decrypting: %s", err)
	}

	// Check that the decrypted plaintext matches the original plaintext
	if string(decrypted) != string(plaintext) {
		t.Errorf("Decrypted plaintext does not match original plaintext: %s", decrypted)
	}
}
