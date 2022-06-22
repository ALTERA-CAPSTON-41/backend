package configs

import "github.com/spf13/viper"

type ServerConfig struct {
	ConnectionString      string `mapstructure:"DB_DSN"`
	MongoConnectionString string `mapstructure:"MONGO_DSN"`
	JWTsecret             string `mapstructure:"JWT_SECRET"`
}

// `path` variable stands for env location, which `.`
// is indicated for env file was located in a same path
// with main.go file.
//
// If you will made an integration test, which is they're
// centralized, perhaps you made a testing environment
// which is located at same folder with tests' folder. So,
// you can argue to the `path` argument with `tests` value.
// instead of `.`
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
