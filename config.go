package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var configInstance = new(Config)

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

func readConfig() (c *Config, err error) {
	data, err := ioutil.ReadFile(configFilename())
	if err != nil {
		c = new(Config)
	} else {
		err = yaml.Unmarshal([]byte(data), &c)
	}
	return
}

func (c *Config) save() (err error) {
	fmt.Printf("%v\n", c)
	d, err := yaml.Marshal(c)
	fmt.Printf("%v\n", c)
	err = ioutil.WriteFile(configFilename(), d, 0600)
	check(err)
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
