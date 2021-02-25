package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	p "swapi/models"

	_ "github.com/lib/pq"
)

// PostgreSQL database info
const (
	host     = "localhost"
	port     = 5432
	user     = "rajaram"
	password = "rajaram"
	dbname   = "practice_db"
)

// Model is Interface has DB functions
type Model interface {
	CreateTable(*sql.DB)
	Insert(int, *sql.DB)
	Delete(int, *sql.DB)
	Update(string, int, *sql.DB)
	Drop(*sql.DB)
}

func main() {

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening a connection to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("DB Successfully connected!")
	var person Model = p.People{}
	// testing purposes
	person.Drop(db)

	person.CreateTable(db)

	fetchAndAdd(1, db)
	fetchAndAdd(2, db)
	fetchAndAdd(3, db)

	person.Update("RAJARAM", 2, db)
	person.Update("JOSHI", 1, db)

	person.Delete(1, db)
	person.Delete(3, db)

}

// This will fetch people from SWAPI api and add to DB
func fetchAndAdd(id int, Db *sql.DB) {
	const url = "https://swapi.dev/api/people/"
	resp, err := http.Get(url + strconv.Itoa(id) + "/")
	log.Println("Fetching from " + url + strconv.Itoa(id) + "/")
	if err != nil {
		log.Fatal("Error", err)
	}
	r := p.People{}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
	r.Insert(id, Db)
}
