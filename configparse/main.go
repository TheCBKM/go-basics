// 1. if no config file it will panic
// 2. if invalid file syntax it will panic
// 3. if wrong file path it will panic
// 4. if no field mention it will set to defaults
// 5. if multiple config files priority be like json > yml > env
// 6. if Unable to deqode the struct it will panic

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
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

var defaults = map[string]interface{}{
	"grpc_port":             "3000",
	"http_port":             "3000",
	"graphql_port":          "3000",
	"datastore_db_host":     "127.0.0.2",
	"datastore_db_user":     "datastore_db_user",
	"datastore_db_password": "datastore_db_password",
	"datastore_db_schema":   "datastore_db_schema",
	"log_level":             10,
	"log_time_format":       "hh:mm:ss",
}

func setConf(configuration *Config, path string) {
	file, err := ReadFile(path)
	extension := filepath.Ext(path)

	if err != nil {
		log.Panicf("unable to read file, %v", err)
	}
	var c map[string]interface{}
	if extension == ".json" {
		err = json.Unmarshal(file, &c)
	} else if extension == ".yml" {
		err = yaml.Unmarshal(file, &c)
	} else if extension == ".env" {
		c = ParseENV(string(file))
	}
	if err != nil {
		log.Panicf("Unable to decode into struct, %v", err)
	}

	var defaultc map[string]interface{}
	err = viper.Unmarshal(&defaultc)
	if err != nil {
		log.Panicf("Unable to decode into struct, %v", err)
	}
	parse(configuration, defaultc)
	parse(configuration, c)
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
	for key := range defaults {
		viper.SetDefault(key, defaults[key])
	}
}

func ParseENV(file string) map[string]interface{} {
	output := map[string]interface{}{}
	for _, line := range strings.Split(file, "\n") {
		data := strings.Split(line, "=")
		if len(data) == 2 {
			data[1]=strings.Replace(data[1],"\"","",-1)
			output[strings.ToLower(strings.Trim(data[0], "\""))] = strings.TrimSpace(data[1])
		}
	}
	return output
}

func CheckExist(path string) (bool, []string, error) {
	str, err := filepath.Glob(path)
	if len(str) > 0 {
		return true, str, err
	} else {
		return false, str, err
	}
}
func ReadFile(path string) ([]byte, error) {
	existed, _, _ := CheckExist(path)
	if existed {
		dat, err := ioutil.ReadFile(path)
		return dat, err
	} else {
		return nil, errors.New("file not exists")
	}
}

func main() {
	setDefaults()
	conf := Config{}
	setConf(&conf, "./configs/.env")
	log.Println(conf)
}
