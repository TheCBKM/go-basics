// 1. if no config file it will panic
// 2. if invalid file syntax it will panic
// 3. if wrong file path it will panic
// 4. if no field mention it will set to defaults
// 5. if multiple config files priority be like json > yml > env
// 6. if Unable to deqode the struct it will panic

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

func setConf(name, path string) conf {
	//call to set default values
	setDefaults()
	//set file path and name
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	//check for file
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}
	//deqode to conf struct
	var c conf
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Panicf("Unable to decode into struct, %v", err)
	}
	log.Println(c)
	return c
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
