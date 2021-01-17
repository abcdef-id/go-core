package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	//driver
	_ "github.com/mattn/go-sqlite3"
)

var (
	// AppPath application path
	AppPath string
)

func init() {
	// set config based on env
	loadEnvVars()
	RedisConnect()
	MongoConnect()
	OpenDbPool()
}

func loadEnvVars() {
	// Bind OS environment variable
	viper.SetEnvPrefix("app")
	viper.BindEnv("env")

	dir, _ := os.Getwd()
	AppPath := dir
	if viper.Get("env") == "testing" {
		viper.SetConfigName("testing")
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("json")
	viper.AddConfigPath(AppPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
