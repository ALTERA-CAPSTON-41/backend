package configs

import "github.com/spf13/viper"

type ServerConfig struct {
	ConnectionString string `mapstructure:"DB_DSN"`
	JWTsecret        string `mapstructure:"JWT_SECRET"`
}

func LoadServerConfig(path string) (ServerConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	config := ServerConfig{}
	err := viper.ReadInConfig()
	viper.Unmarshal(&config)
	return config, err
}
