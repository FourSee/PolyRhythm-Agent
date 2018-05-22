package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var configInstance = readConfig()

type Config struct {
	PairedDevice struct {
		PublicKey string
	}
	DeviceIdentity struct {
		PublicKey  string
		PrivateKey string
	}
}

func config() *Config {
	return configInstance
}

func readConfig() (c *Config) {
	c = new(Config)
	data, err := ioutil.ReadFile(configFilename())
	if err == nil {
		err = yaml.Unmarshal([]byte(data), &c)
		check(err)
	}
	return
}

func (c *Config) save() (err error) {
	d, err := yaml.Marshal(c)
	err = ioutil.WriteFile(configFilename(), d, 0600)
	check(err)
	return
}

func configFilename() string {
	return filepath.Join(configDir(), "config.yaml")
}

func configDir() string {
	configEnv, p := os.LookupEnv("POLYRHYTHM_CONFIG_DIR")
	if p {
		return configEnv
	}
	home, _ := homedir.Dir()
	path := filepath.Join(home, ".polyrhytm")
	err := os.MkdirAll(path, os.ModePerm)
	check(err)
	return path
}

func check(e error) {
	if e != nil {
		fmt.Printf("PolyRhythm had a fatal error: %v\n", e)
		os.Exit(0)
	}
}
