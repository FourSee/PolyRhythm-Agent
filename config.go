package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var configInstance = readConfig()

// Config is the global configuration structure
type Config struct {
	PairedDevice struct {
		PublicKey string `yaml:"PublicKey"`
		ID        string `yaml:"ID"`
	} `yaml:"PairedDevice"`
	DeviceIdentity struct {
		PublicKey  string `yaml:"PublicKey"`
		PrivateKey string `yaml:"PrivateKey"`
		ID         string `yaml:"ID"`
	} `yaml:"DeviceIdentity"`
}

func config() *Config {
	return configInstance
}

func readConfig() (c *Config) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	fmt.Println("Looking for config.yaml in...")
	viper.AddConfigPath("/etc/polyrhytm/") // path to look for the config file in
	fmt.Println("/etc/polyrhytm/")
	viper.AddConfigPath(homeConfigDir()) // call multiple times to add many search paths
	fmt.Println(homeConfigDir())
	viper.AddConfigPath(workingPath()) // optionally look for config in the working directory
	fmt.Println(workingPath())
	err := viper.ReadInConfig() // Find and read the config file

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		mapstructure.Decode(viper.AllSettings(), c)
	})
	if err != nil {
		fmt.Println("No valid config files found. If you're pairing a device, ignore this")
		c = new(Config)
	} else {
		err = mapstructure.Decode(viper.AllSettings(), &c)
		if err != nil {
			fmt.Printf("Can't read %s: %s\n", viper.ConfigFileUsed(), err)
		}
	}
	return
}

func (c *Config) save() {
	d, err := yaml.Marshal(c)
	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		configFile, _ = filepath.Abs("./config.yaml")
	}
	fmt.Printf("Writing to %s\n", configFile)
	err = ioutil.WriteFile(configFile, d, 0600)
	check(err)
}

func homeConfigDir() (d string) {
	d, _ = homedir.Expand("~/.polyrhythm/")
	return
}

func workingPath() (d string) {
	d, _ = filepath.Abs("./")
	return
}

func check(e error) {
	if e != nil {
		fmt.Printf("PolyRhythm had a fatal error: %v\n", e)
		os.Exit(0)
	}
}
