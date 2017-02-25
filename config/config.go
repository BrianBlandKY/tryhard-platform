package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
    "os"
)

type CSql struct {
	Username string `yaml:"user"`
	Password string `yaml:"pwd"`
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
}
type CFileSystem struct {
	Directory  	string      `yaml:"directory"`
	Permission 	os.FileMode `yaml:"permission"`
}
type CStore struct {
	Sql        CSql        `yaml:"sql"`
	FileSystem CFileSystem `yaml:"file_system"`
}
type CApi struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}
type Config struct {
    Env   string   `yaml:"env"`
	Store CStore   `yaml:"store"`
	Api   CApi     `yaml:"api"`
}

func ParseConfig(path string) (config Config) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config = Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	return config
}