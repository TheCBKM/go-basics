package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	setConf("config", "./configs")
}

func TestPanicJSON(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	createFile("./configs/config.json")

	var file, err = os.OpenFile("./configs/config.json", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(`{
		"host"     : "fromjson",
		"port"     : 1234,
		"user"     : "rajaram",
		"password" : [1,2,3]
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
	setConf("config", "./configs")
	deleteFile("./configs/config.json")

}
func TestJSON(t *testing.T) {
	createFile("./configs/config.json")
	createJSON()
	c := setConf("config", "./configs")
	if (conf{Host: "fromjson", Port: 1234, User: "rajaram", Password: "rajaram", Dbname: "default_db"} != c) {
		t.Error("WRONG ")
	}
	log.Println(c)

	deleteFile("./configs/config.json")
}

func TestYML(t *testing.T) {
	createFile("./configs/config.yml")
	createYML()
	c := setConf("config", "./configs")
	if (conf{Host: "fromyml", Port: 8080, User: "sss", Password: "default_pass", Dbname: "practice_db2"} != c) {
		t.Error("WRONG ")
	}
	log.Println(c)

	deleteFile("./configs/config.yml")
}

func TestENV(t *testing.T) {
	createFile("./configs/config.env")
	createENV()
	c := setConf("config", "./configs")
	if (conf{Host: "fromenv", Port: 1111, User: "default_user", Password: "default_pass", Dbname: "default_db"} != c) {
		t.Error("WRONG ")
	}
	log.Println(c)

	deleteFile("./configs/config.env")
}

func createJSON() {
	var file, err = os.OpenFile("./configs/config.json", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(`{
		"host"     : "fromjson",
		"port"     : 1234,
		"user"     : "rajaram",
		"password" : "rajaram"
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

func createYML() {
	var file, err = os.OpenFile("./configs/config.yml", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("host: fromyml \nuser: sss \ndbname: practice_db2")
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
	var file, err = os.OpenFile("./configs/config.env", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("HOST=fromenv\nPORT=1111")
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
