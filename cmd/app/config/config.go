package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBHOST     string `mapstructure:"PGHOST"`
	DBUSER     string `mapstructure:"PGUSER"`
	DBPASSWORD string `mapstructure:"PGPASSWORD"`
	DBNAME     string `mapstructure:"PGDB"`
	ADDRESS    string `mapstructure:"ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(ctx context.Context, path string, logger *logrus.Logger) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	return
}
