package config

import (
	"base/server/logging"
	"fmt"
	"os"

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
	Host     string
	Port     string
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
		panic(fmt.Sprintf("Fatal error reading config file: %s \n", err))
	}

	dbHost := os.Getenv(viper.GetString(DbHost))
	dbPassword := os.Getenv(viper.GetString(DbPassword))
	dbUser := os.Getenv(viper.GetString(DbUser))

	if len(dbHost) == 0 || len(dbPassword) == 0 || len(dbUser) == 0 {
		logging.Log("Fatal error, not all database environment variables are set\n")
		panic("Fatal error, not all database environment variables are set\n")
	}

	return Config{
		Server: newServerConfig(
			viper.GetString(ServerHost),
			viper.GetInt(ServerPort),
			viper.GetString(ServerVersion),
		),
		Database: newDatabaseConfig(
			dbHost,
			viper.GetString(DbPort),
			viper.GetString(DbName),
			dbUser,
			dbPassword,
		),
	}
}

func newServerConfig(host string,
	port int,
	version string,
) serverConfig { // private
	return serverConfig{
		Host:    host,
		Port:    port,
		Version: version,
	}
}

// TO DO
func newDatabaseConfig(host string,
	port string,
	name string,
	user string,
	password string,
) databaseConfig { // private

	return databaseConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: password,
	}
}
