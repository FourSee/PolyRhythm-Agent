package main

import (
	b64 "encoding/base64"
	"testing"

	"golang.org/x/crypto/nacl/box"
)

func TestGenerateCryptoKey(t *testing.T) {
	_, _, err := generateCryptoKey()
	if err != nil {
		t.Errorf("Got an error generating keys: %s\n", err)
	}
}

func BenchmarkGenerateCryptoKey(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := generateCryptoKey()
		if err != nil {
			panic(err)
		}
	}
}

func TestEncryptString(t *testing.T) {
	alicePubKey, alicePrivKey, err := generateCryptoKey()
	if err != nil {
		t.Errorf("Got an error generating keys: %s\n", err)
	}

	bobPubKey, bobPrivKey, err := generateCryptoKey()
	if err != nil {
		t.Errorf("Got an error generating keys: %s\n", err)
	}

	msg := "This is a simple test message"

	alicePubBytes := KeyToBytes(alicePubKey)
	bobPrivBytes := KeyToBytes(bobPrivKey)

	encryptedMsg := EncryptString(msg, &alicePubBytes, &bobPrivBytes)

	encrypted, err := b64.StdEncoding.DecodeString(*encryptedMsg)
	if err != nil {
		t.Errorf("Decryption error: %s\n", err)
	}
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	alicePrivBytes := KeyToBytes(alicePrivKey)
	bobPubBytes := KeyToBytes(bobPubKey)

	decrypted, ok := box.Open(nil, encrypted[24:], &decryptNonce, &bobPubBytes, &alicePrivBytes)
	if !ok {
		t.Error("Decryption error\n")
	}

	if string(decrypted) != msg {
		t.Error("Decryption error: messages do not match\n")
	}

}

func BenchmarkEncrypt(b *testing.B) {
	b.ResetTimer()
	alicePubKey, _, err := generateCryptoKey()
	if err != nil {
		b.Errorf("Got an error generating keys: %s\n", err)
	}

	_, bobPrivKey, err := generateCryptoKey()
	if err != nil {
		b.Errorf("Got an error generating keys: %s\n", err)
	}

	msg := "This is a simple test message"
	alicePubBytes := KeyToBytes(alicePubKey)
	bobPrivBytes := KeyToBytes(bobPrivKey)

	for i := 0; i < b.N; i++ {
		EncryptString(msg, &alicePubBytes, &bobPrivBytes)
	}
}
