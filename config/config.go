package config

import (
	"errors"
	"os"
	"serv/pkg/logger"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var (
	Instance Config
)

type Config struct {
	Server
	Database
	Cache
}

type Server struct {
	Addr string `yaml:"addr"`
}

type Database struct {
	Addr string `yaml:"addr"`
}

type Cache struct {
	Time int64 `yaml:"time"`
}

type rootConfig struct {
	Server   `yaml:"server"`
	Database `yaml:"database"`
	Cache    `yaml:"cache"`
}

func NewConfigYaml(path string) (*Config, error) {
	rootConfig := &rootConfig{}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(rootConfig)

	if err != nil {
		return nil, err
	}

	return &Config{
		Server:   rootConfig.Server,
		Database: rootConfig.Database,
		Cache:    rootConfig.Cache,
	}, nil

}

func NewConfigEnv() (*Config, error) {

	for _, v := range []string{"srvAddr", "dbAddr", "cachetime"} {

		if _, set := os.LookupEnv(v); !set {
			return nil, errors.New("Err of environment settings")
		}
	}

	return &Config{
		Server: Server{
			Addr: os.Getenv("srvAddr"),
		},
		Database: Database{
			Addr: os.Getenv("dbAddr"),
		},
		Cache: Cache{
			Time: getEnvAsInt(os.Getenv("time"), 0),
		},
	}, nil

}

func NewConfigDotEnv(envConfigPath string) (*Config, error) {

	if err := godotenv.Load(envConfigPath); err != nil {
		return nil, err
	}

	return NewConfigEnv()
}

func getEnvAsInt(str string, def int64) int64 {
	if val, err := strconv.Atoi(str); err == nil {
		return int64(val)
	}
	return def
}

func InitConfig() {
	configYamlpath := "D:/Golang/service/confi.yml"
	configDotEnvpath := "D:/Golang/service/.env"

	inst, err := NewConfigYaml(configYamlpath)
	if err != nil {
		inst, err = NewConfigDotEnv(configDotEnvpath)
		if err != nil {
			inst, err = NewConfigEnv()
			if err != nil {
				logger.Err("Bad config data", err)
			}
		}
	}
	Instance = *inst
}
