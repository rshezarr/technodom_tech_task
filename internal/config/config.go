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

	//unmarshal config into "config" struct
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error while parsing configs: %v", err)
		// return nil
	}

	return cfg
}

func InitConfig() error {
	//set flag
	var configPath = flag.String("config-path", "configs/", "path to config file")

	//parse flag
	flag.Parse()

	//set config file path, name and type
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configPath)

	//read config by property above
	return viper.ReadInConfig()
}
