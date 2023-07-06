package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	// "strconv"
	"sync"
)

type AppConfig struct {
	Env              string `mapstructure:"env"`
	ServerHost       string `mapstructure:"server_host"`
	ServerPort       string `mapstructure:"server_port"`
	MySQLURI         string `mapstructure:"mysql_uri"`
}

var (
	instance *AppConfig = nil
	once     sync.Once
)

func getEnv(key string, fallback interface{}) interface{} {
	var rValue interface{}
	value, exists := os.LookupEnv(key)
	if !exists {
		rValue = fallback
	} else {
		rValue = value
	}
	return rValue
}

func InitConfig() *AppConfig {
	// synchronize instance config
	once.Do(
		func() {
			var err error
			// load .env config
			err = godotenv.Load(".env")
			if err != nil {
				log.Println("Error: ", err)
			}
			if err != nil {
				log.Println("Error: ", err)
			}
			instance = &AppConfig{
				Env:              getEnv("ENV", "development").(string),
				ServerHost:       getEnv("SERVER_HOST", "0.0.0.0").(string),
				ServerPort:       getEnv("SERVER_PORT", "8088").(string),
				MySQLURI:         getEnv("MYSQL_URI", "username:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local").(string),
			}
		},
	)

	return instance
}
