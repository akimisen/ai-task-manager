package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RabbitMQ RabbitMQConfig
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

type ServerConfig struct {
	Address string `mapstructure:"address"`
}

type DatabaseConfig struct {
	MySQL MySQLConfig `mapstructure:"mysql"`
	Redis RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Addr string `mapstructure:"addr"`
}

type AuthConfig struct {
	JWT JWTConfig `mapstructure:"jwt"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")                      // 当前目录
	viper.AddConfigPath("./configs")              // configs 子目录
	viper.AddConfigPath("$HOME/.ai-task-manager") // 用户主目录
	viper.AddConfigPath("/etc/ai-task-manager/")  // 系统级配置目录

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
