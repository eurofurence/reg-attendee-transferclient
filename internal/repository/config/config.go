package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	configurationData     *conf
)

type regsysConfig struct {
	Url   string `yaml:"url"`
	Token string `yaml:"token"`
}

type attendeeSrvConfig struct {
	Url string `yaml:"url"`
}

type conf struct {
	Regsys      regsysConfig      `yaml:"regsys"`
	AttendeeSrv attendeeSrvConfig `yaml:"attendee-service"`
}

func LoadConfiguration() {
	configurationFilename := "config.yaml"
	log.Printf("INFO  loading configuration from %v", configurationFilename)

	yamlFile, err := ioutil.ReadFile(configurationFilename)
	if err != nil {
		log.Fatalf("FATAL failed to load configuration file '%s': %v", configurationFilename, err)
	}
	newConfigurationData := &conf{}
	err = yaml.UnmarshalStrict(yamlFile, newConfigurationData)
	if err != nil {
		log.Fatalf("FATAL failed to parse configuration file '%s': %v", configurationFilename, err)
	}
	configurationData = newConfigurationData

	log.Print("INFO  success")
}

func RegsysBaseUrl() string {
	return configurationData.Regsys.Url
}

func AttendeeServiceBaseUrl() string {
	return configurationData.AttendeeSrv.Url
}

func RegsysTransferToken() string {
	return configurationData.Regsys.Token
}
