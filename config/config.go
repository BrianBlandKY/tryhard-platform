// Configuration Abstraction for YAML.
// Other solutions should not reference YAML directly

package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config Default Confuguration Model
type Config struct {
	Settings struct {
		VerboseHeartbeat bool `yaml:"verbose_heartbeat"`
		VerboseMessaging bool `yaml:"verbose_messaging"`
		ConsoleEnabled   bool `yaml:"console_enabled"`
	}
	Service struct {
		Name string `yaml:"name"`
		ID   string `yaml:"id"`
	}
	Platform struct {
		Address string `yaml:"address"`
	}
}

// ReadConfig
func ReadConfig(file string) Config {
	filename, _ := filepath.Abs(file)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return cfg
}

// WriteConfig
func WriteConfig(cfg Config) string {
	output, _ := yaml.Marshal(cfg)
	return string(output[:])
}
