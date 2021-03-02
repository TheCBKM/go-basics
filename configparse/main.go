// 1. if no config file it will panic
// 2. if invalid file syntax it will panic
// 3. if wrong file path it will panic
// 4. if no field mention it will set to defaults
// 5. if multiple config files priority be like json > yml > env
// 6. if Unable to deqode the struct it will panic

package main

import (
	"log"
	"reflect"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GrpcPort string `yaml:"grpc_port" json:"grpc_port"`
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HttpPort string `yaml:"http_port" json:"http_port"`
	//Graphql Port
	GraphqlPort string `yaml:"graphql_port" json:"graphql_port"`
	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDbHost string `yaml:"datastore_db_host" json:"datastore_db_host"`
	// DatastoreDBUser is username to connect to database
	DatastoreDbUser string `yaml:"datastore_db_user" json:"datastore_db_user"`
	// DatastoreDBPassword password to connect to database
	DatastoreDbPassword string `yaml:"datastore_db_password" json:"datastore_db_password"`
	// DatastoreDBSchema is schema of database
	DatastoreDbSchema string `yaml:"datastore_db_schema" json:"datastore_db_schema"`
	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int `yaml:"log_level" json:"log_level"`
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string `yaml:"log_time_format" json:"log_time_format"`
}

const defaultJSON = `{
    "grpc_port": "8082",
    "http_port": "1234",
    "graphql_port": "1111",
    "datastore_db_host": "127.0.0.2",
    "datastore_db_user": "rajaram",
    "datastore_db_password": "pas123",
    "datastore_db_schema": "schema",
    "log_level": 1,
    "log_time_format": "hh:mm:ss"
}`

func setConf(name, path string, env bool) {
	//call to set default values
	// setDefaults()
	if env {
		viper.SetConfigFile(".env")

	} else {
		//set file path and name
		viper.SetConfigName(name)
		viper.AddConfigPath(path)

	}

	//check for file
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}
	//deqode to Config struct
	var c map[string]interface{}
	var configuration Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Panicf("Unable to decode into struct, %v", err)
	}
	parse(&configuration, c)
	log.Println(configuration)
	log.Println(viper.Get("grpc_port"))

}

func parse(configuration *Config, resp map[string]interface{}) *Config {
	result := map[string]interface{}{}
	t := reflect.TypeOf(*configuration)
	for i := 0; i < 9; i++ {
		field := t.Field(i)
		if resp[field.Tag.Get("json")] != nil {
			if field.Tag.Get("json") == "log_level" && reflect.TypeOf(resp[field.Tag.Get("json")]) == reflect.TypeOf("string") {
				i, _ := strconv.Atoi(resp[field.Tag.Get("json")].(string))
				result[field.Name] = i
			} else {

				result[field.Name] = resp[field.Tag.Get("json")]

			}
		}
	}
	mapstructure.Decode(result, &configuration)
	return configuration

}

func setDefaults() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("user", "default_user")
	viper.SetDefault("host", "default_host")
	viper.SetDefault("password", "default_pass")
	viper.SetDefault("dbname", "default_db")
}
func main() {
	setConf("config", "./configs", false)

}
