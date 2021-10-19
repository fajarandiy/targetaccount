package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
//The values are read by viper from a config file or environment variables.
type Config struct {
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBSource        string `mapstructure:"DB_SOURCE"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	ConnMaxLifetime int    `mapstructure:"CONN_MAX_LIFETIME"`
	MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
