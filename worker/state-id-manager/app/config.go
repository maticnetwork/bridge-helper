package app

import "github.com/spf13/viper"

// Reading configuration params in memory, during
// application boot up
func read(file string) error {
	viper.SetConfigFile(file)

	return viper.ReadInConfig()
}

// Retrieves value for specified key
func get(key string) string {
	return viper.GetString(key)
}
