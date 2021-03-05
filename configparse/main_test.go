package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic")
		}
	}()
	setConf(&Config{}, "config")
}

func TestJSON(t *testing.T) {
	createFile("./test_configs/config.json")
	createJSON()
	c := Config{}
	tc := Config{"8082", "1234", "1111", "127.0.0.2", "json", "pas123", "schema", 0, "hh:mm:ss"}
	setConf(&c, "./test_configs/config.json")
	if tc != c {
		t.Error("wrong in json ")
	}

	deleteFile("./test_configs/config.json")
}

func TestYML(t *testing.T) {
	createFile("./test_configs/config.yml")
	createYML()
	c := Config{}
	tc := Config{"8080", "1234", "1111", "127.0.0.1", "rajaram", "pas123", "schema", 1, "hh:mm:ss"}
	setConf(&c, "./test_configs/config.yml")
	if tc != c {
		t.Error("wrong in yml ")
	}

	deleteFile("./test_configs/config.yml")
}

func TestENV(t *testing.T) {
	createFile("./test_configs/.env")
	createENV()
	c := Config{}
	tc := Config {"8080" ,"1234" ,"1111","127.0.0.1", "ENV" ,"pas123" ,"schema", 2 ,"hh:mm:ss"}
	setConf(&c, "./test_configs/.env")
	if tc != c {
		t.Error("wrong in yml ")
	}

	deleteFile("./test_configs/.env")
}

func createJSON() {
	var file, err = os.OpenFile("./test_configs/config.json", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(`{
    "grpc_port": "8082",
    "http_port": "1234",
    "graphql_port": "1111",
    "datastore_db_host": "127.0.0.2",
    "datastore_db_user": "json",
    "datastore_db_password": "pas123",
    "datastore_db_schema": "schema",
    "log_time_format": "hh:mm:ss"
}`)
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

//
func createYML() {
	var file, err = os.OpenFile("./test_configs/config.yml", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("grpc_port: \"8080\"\nhttp_port: \"1234\"\ngraphql_port: \"1111\"\ndatastore_db_host: \"127.0.0.1\"\ndatastore_db_user: \"rajaram\"\ndatastore_db_password: \"pas123\"\ndatastore_db_schema: \"schema\"\nlog_level: 1\nlog_time_format: \"hh:mm:ss\"\n")
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

func createENV() {
	var file, err = os.OpenFile("./test_configs/.env", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("grpc_port= \"8080\"\nhttp_port= \"1234\"\ngraphql_port= \"1111\"\ndatastore_db_host= \"127.0.0.1\"\ndatastore_db_user= \"ENV\"\ndatastore_db_password= \"pas123\"\ndatastore_db_schema= \"schema\"\nlog_level= 2\nlog_time_format= \"hh:mm:ss\"\n")
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}
