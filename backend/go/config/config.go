package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	BackendConfig `mapstructure:"backend"`
	MongodbConfig `mapstructure:"mongo"`
	// JwtSecret     string
	// Env           string
}

type BackendConfig struct {
	Port int
}

type MongodbConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Database   string
	Collection string
}

func Load() *Config {
	viper.AddConfigPath("/etc/k8s-todos")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal().Msg("No configuration file found")
		} else {
			log.Info().Msg("Use configuration file")
		}
	}

	viper.SetConfigName("secret")
	viper.SetConfigType("toml")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal().Msg("No secret file found")
		} else {
			log.Info().Msg("Use configuration file")
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal().Err(err)
	}

	return &config
}

func readConfigFile(fileName, fileType, filePath, errMsgIfNotFound string) {
	viper.SetConfigFile(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal().Msg("No configuration file found")
		} else {
			log.Info().Msg("Use configuration file")
		}
	}
}
