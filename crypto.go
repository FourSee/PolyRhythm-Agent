package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	crypto "github.com/foursee/shellgameCrypto"
)

func myCryptoKey() (privKey, pubKey string, err error) {
	if configInstance.DeviceIdentity.PublicKey == "" && configInstance.DeviceIdentity.PrivateKey == "" {
		return generateCryptoKey()
	}
	return configInstance.DeviceIdentity.PrivateKey, configInstance.DeviceIdentity.PublicKey, nil
}

func generateCryptoKey() (string, string, error) {
	log.Output(0, fmt.Sprintf("Generating %v-bit RSA keypair...", *keyBits))
	user, _ := user.Current()
	name := user.Username
	comment := "PolyRythm generated keypair"
	hostname, _ := os.Hostname()
	email := fmt.Sprintf("%s@%s", name, hostname)
	privKey, pubKey, err := crypto.GenerateRSAKeyPair(*keyBits, name, comment, email)
	check(err)
	configInstance.DeviceIdentity.PrivateKey = privKey
	configInstance.DeviceIdentity.PublicKey = pubKey
	configInstance.save()
	return privKey, pubKey, err

}
