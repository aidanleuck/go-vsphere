package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type dbConfig struct {
	DBPort           uint32
	ConnectionString string
}

type appConfiguration struct {
	DB   dbConfig
	Port uint32
}

var appConfigMap *appConfiguration

func GetConfiguration() *appConfiguration {
	if appConfigMap == nil {
		LoadConfiguration()
	}
	return appConfigMap
}

func LoadConfiguration() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name // call multiple times to add many search paths

	rootPath := path.Dir(basepath) // optionally look for config in the working directory
	viper.AddConfigPath(rootPath)

	viper.SetEnvPrefix("VSPHERE_AGENT")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("PORT")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	dbPortConfig := viper.GetUint32("DB_PORT")
	dbHost := viper.GetString("DB_HOST")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	agentPort := viper.GetUint32("PORT")

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbUser, dbPassword, dbHost, dbPortConfig)

	dbStruct := dbConfig{
		DBPort:           dbPortConfig,
		ConnectionString: connectionString,
	}

	appConfigStruct := appConfiguration{
		DB:   dbStruct,
		Port: agentPort,
	}

	appConfigMap = &appConfigStruct
}
