package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	API struct {
		Port           string `yaml:"port"`
		Host           string `yaml:"host"`
		MaxHeaderBytes int    `yaml:"maxHeaderBytes"`
		Timeout        int    `yaml:"timeout"`
	} `yaml:"api"`
}

func NewConfig() *Config {
	var cfg *Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error while parsing configs: %v", err)
		// return nil
	}

	return cfg
}

func InitConfig() error {
	var configPath = flag.String("config-path", "configs/", "path to config file")

	flag.Parse()

	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configPath)

	return viper.ReadInConfig()
}
