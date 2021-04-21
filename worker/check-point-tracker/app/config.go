package app

import "github.com/spf13/viper"

// Reading configuration into memory
func read(file string) error {
	viper.SetConfigFile(file)

	return viper.ReadInConfig()
}

// Retrieving value for given key
func get(key string) string {
	return viper.GetString(key)
}
