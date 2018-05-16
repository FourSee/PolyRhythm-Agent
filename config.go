package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PairedDevice struct {
		PublicKey string
	}
	DeviceIdentity struct {
		PublicKey  string
		PrivateKey string
	}
}

func getConfig() (c Config) {
	c, _ = readConfig()
	return c
}

func readConfig() (c Config, err error) {
	data, err := ioutil.ReadFile(configFilename())
	err = yaml.Unmarshal([]byte(data), &c)
	return
}

func writeConfig(c Config) (err error) {
	d, err := yaml.Marshal(&c)
	err = ioutil.WriteFile(configFilename(), d, 0600)
	return
}

func configFilename() string {
	return filepath.Join(configDir(), "config.yaml")
}

func configDir() string {
	home, _ := homedir.Dir()
	path := filepath.Join(home, ".polyrhytm")
	err := os.MkdirAll(path, os.ModePerm)
	check(err)
	return path
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
