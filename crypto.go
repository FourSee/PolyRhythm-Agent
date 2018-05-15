package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	crypto "github.com/foursee/shellgameCrypto"
)

func myCryptoKey() (privKey, pubKey string, err error) {
	return generateCryptoKey()
}

func generateCryptoKey() (string, string, error) {
	log.Output(0, "Generating 2048-bit RSA keypair...")
	user, _ := user.Current()
	name := user.Username
	comment := "PolyRythm generated keypair"
	hostname, _ := os.Hostname()
	email := fmt.Sprintf("%s@%s", name, hostname)
	return crypto.GenerateRSAKeyPair(2048, name, comment, email)
}
