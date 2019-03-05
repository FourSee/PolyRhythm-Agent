package main

import (
	"bytes"
	crypto_rand "crypto/rand"
	b64 "encoding/base64"
	"io"

	"golang.org/x/crypto/nacl/box"
)

func myCryptoKey() (privKey, pubKey string, err error) {
	if config().DeviceIdentity.PublicKey == "" || config().DeviceIdentity.PrivateKey == "" {

		config().DeviceIdentity.PrivateKey, config().DeviceIdentity.PublicKey, err = generateCryptoKey()
		check(err)
		config().save()
	}
	return config().DeviceIdentity.PrivateKey, config().DeviceIdentity.PublicKey, nil
}

func generateCryptoKey() (pub string, priv string, err error) {
	// log.Output(0, "Generating ECDSA keypair...")
	pubKey, privKey, err := box.GenerateKey(crypto_rand.Reader)
	if err != nil {
		return
	}
	pub = b64.StdEncoding.EncodeToString(pubKey[:])
	priv = b64.StdEncoding.EncodeToString(privKey[:])

	// user, _ := user.Current()
	// name := user.Username
	// comment := "PolyRythm generated keypair"
	// hostname, _ := os.Hostname()
	// email := fmt.Sprintf("%s@%s", name, hostname)
	// privKey, pubKey, err := crypto.GenerateRSAKeyPair(keyBits(), name, comment, email)
	// check(err)
	return
}

func nonce() (n [24]byte) {
	if _, err := io.ReadFull(crypto_rand.Reader, n[:]); err != nil {
		panic(err)
	}
	return
}

// EncryptReader encrypts a byte array to a base64-encoded nacl box
func EncryptReader(r io.Reader, pub, sec *[32]byte) (encrypted *string) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	check(err)
	return EncryptBytes(buf.Bytes(), pub, sec)
}

// EncryptString encrypts a byte array to a base64-encoded nacl box
func EncryptString(msg string, pub, sec *[32]byte) (encrypted *string) {
	return EncryptBytes([]byte(msg), pub, sec)
}

// EncryptBytes encrypts a byte array to a base64-encoded nacl box
func EncryptBytes(msg []byte, pub, sec *[32]byte) (encrypted *string) {
	n := nonce()
	// alicePub := keyBytes(config().PairedDevice.PublicKey)
	// bobSec := keyBytes(config().DeviceIdentity.PrivateKey)
	encryptedBytes := box.Seal(n[:], msg, &n, pub, sec)
	s := b64.StdEncoding.EncodeToString(encryptedBytes)
	return &s
}

// KeyToBytes converts a string to byte arrays suitable for NaCL use
func KeyToBytes(s string) (b [32]byte) {
	slice, err := b64.StdEncoding.DecodeString(s)
	check(err)
	copy(b[:], slice)
	return
}
