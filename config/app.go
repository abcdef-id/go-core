package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	//driver
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// set config based on env
	LoadEnvVars()
	RedisConnect()
	MongoConnect()
	OpenDbPool()
}

// LoadEnvVars ..
func LoadEnvVars() {
	// Bind OS environment variable
	viper.SetEnvPrefix("app")
	viper.BindEnv("env")

	dir, _ := os.Getwd()
	appPath := dir
	if viper.Get("env") == "testing" {
		viper.BindEnv("path") // bind OS environment variable APP_PATH variable
		viper.SetConfigName("testing")
		appPath = viper.GetString("path")
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("json")
	viper.AddConfigPath(appPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
