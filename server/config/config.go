package config

import (
	"base/server/logging"
	"fmt"

	"github.com/spf13/viper"
)

/**
 * Responsible for mapping values in config.yaml to structs/objects
**/

type Config struct {
	Server   serverConfig
	Database databaseConfig
}

type serverConfig struct {
	Host    string
	Port    int
	Version string
}

type databaseConfig struct {
	Name     string
	User     string
	Password string
}

func Init() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()

	if err != nil {
		logging.Log(fmt.Sprintf("Fatal error reading config file: %s \n", err))
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	return Config{
		Server: newServerConfig(
			viper.GetString(ServerHost),
			viper.GetInt(ServerPort),
			viper.GetString(ServerVersion),
		),
		Database: newDatabaseConfig( // TO DO
			"",
			"",
			"",
		),
	}
}

func newServerConfig(host string, port int, version string) serverConfig { // private
	return serverConfig{
		Host:    host,
		Port:    port,
		Version: version,
	}
}

// TO DO
func newDatabaseConfig(name string, user string, password string) databaseConfig { // private
	return databaseConfig{
		Name:     name,
		User:     user,
		Password: password,
	}
}
