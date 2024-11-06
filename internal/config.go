package internal

import "github.com/spf13/viper"

type Config struct {
	Server    ServerConfig `mapstructure:"server"`
	WebClient WebClient    `mapstructure:"web_client"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type WebClient struct {
	SoapApi SoapApi `mapstructure:"soap_api"`
	RestApi RestApi `mapstructure:"rest_api"`
}

type SoapApi struct {
	Endpoint string    `mapstructure:"endpoint"`
	Paths    SoapPaths `mapstructure:"paths"`
}

type SoapPaths struct {
	UserUrl string `mapstructure:"user_url"`
}

type RestApi struct {
	Endpoint string    `mapstructure:"endpoint"`
	Paths    RestPaths `mapstructure:"paths"`
}

type RestPaths struct {
	UserUrl string `mapstructure:"user_url"`
}

func LoadConfig() (*Config, error) {
	var conf Config

	// Set file name and path
	viper.SetConfigName("config")   // Assuming config.yaml
	viper.AddConfigPath("./config") // Current directory

	// Read in the config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal into Config struct
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
