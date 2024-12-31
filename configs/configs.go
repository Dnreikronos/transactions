package configs

import (
	"log"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	DatabasePath string
}

func init() {
	viper.SetDefault("api.port", "9090")
	viper.SetDefault("database.path", "transactions.db")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No configuration file found, using default values")
	}
	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			DatabasePath: viper.GetString("database.path"),
		},
	}
	log.Printf("Configuration loaded: %+v", cfg)
	return nil
}

func GetDB() DBConfig {
	if cfg == nil {
		log.Fatal("Configuration is not initialized")
	}
	return cfg.DB
}

func GetServerPort() string {
	if cfg == nil {
		log.Fatal("Configuration is not initialized")
	}
	return cfg.API.Port
}
