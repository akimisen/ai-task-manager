package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RabbitMQ struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"rabbitmq"`
	Server struct {
		Address string `mapstructure:"address"`
	} `mapstructure:"server"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.tts-service")
	viper.AddConfigPath("/etc/tts-service/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	// 打印配置信息，方便调试
	fmt.Printf("Loaded config: %+v\n", cfg)

	return &cfg
}
