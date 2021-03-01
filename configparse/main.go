package main

import (
	"log"

	"github.com/spf13/viper"
)

type conf struct {
	Host     string `yaml:"host" json:"host""`
	Port     int64  `yaml:"port" json:"port" `
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Dbname   string `yaml:"dbname" json:"dbname"`
}

func setConf(name, path string) {
	setDefaults()
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}
	var c conf
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Panicf("Unable to decode into struct, %v", err)
	}
	log.Println(c)
}

func setDefaults() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("user", "default_user")
	viper.SetDefault("host", "default_host")
	viper.SetDefault("password", "default_pass")
	viper.SetDefault("dbname", "default_db")

}
func main() {
	setConf("config", "./configs")
}
