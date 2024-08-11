package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type MgoConfig struct {
	MongoUri        string `mapstructure:"mongo_uri,omitempty"`
	MongoCollection string `mapstructure:"mongo_collection,omitempty"`
	MongoDb         string `mapstructure:"mongo_db,omitempty"`
}

type Config struct {
	ApiKey string    `mapstructure:"api_key,omitempty"`
	Domain string    `mapstructure:"domain,omitempty"`
	Mgo    MgoConfig `mapstructure:"mgo,omitempty"`
}

func LoadConfig() Config {
	var config Config

	viper.AddConfigPath("./conf") // look for config in the working directory
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		return Config{}
	}

	return config
}
