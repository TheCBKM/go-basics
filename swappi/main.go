package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

type People struct {
	ID        int
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	Homeworld string `json:"homeworld"`
	URL       string `json:"url"`
	Created   string `json:"created"`
	Edited    string `json:"edited"`
}

//Db is pointer to DB
var Db *sql.DB

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
	fmt.Println("DB Successfully connected!")
	Db = db
	dropTable()
	createTable()

	fetchAndAdd(4)
	fetchAndAdd(1)
	fetchAndAdd(2)
	fetchAndAdd(3)

	updateName("RAJARAM", 2)
	updateName("JOSHI", 4)

	deletePeople(1)
	deletePeople(3)
}

//Create table according to Struct and Swapi scheema
func createTable() {
	_, err := Db.Query(`
CREATE TABLE IF NOT EXISTS people (
id INT PRIMARY KEY,
Name      TEXT ,
Height    TEXT ,
Mass      TEXT ,
HairColor TEXT ,
SkinColor TEXT ,
EyeColor  TEXT ,
BirthYear TEXT ,
Gender    TEXT ,
Homeworld TEXT ,
URL       TEXT ,
Created   TEXT ,
Edited    TEXT 
	  );
	  `)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Created")
}

// This will fetch people from SWAPI api and add to DB
func fetchAndAdd(id int) {
	resp, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(id) + "/")
	fmt.Println("Fetching from https://swapi.dev/api/people/" + strconv.Itoa(id) + "/")
	if err != nil {
		log.Fatal("Error", err)
	}
	r := People{}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r.Name)
	insertPeople(r, id)
}

//Insert People to DB
func insertPeople(p People, id int) {
	sqlStatement := `
	INSERT INTO people (id,name , height , mass , haircolor , skincolor , eyecolor , birthyear , gender , homeworld , url , created , edited )
	VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	RETURNING id`
	err := Db.QueryRow(sqlStatement, id, p.Name, p.Height, p.Mass, p.HairColor, p.SkinColor, p.EyeColor, p.BirthYear, p.Gender, p.Homeworld, p.URL, p.Created, p.Edited).Scan(&p.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("People Inserted", p.ID)

}

//Delete People on ID
func deletePeople(id int) {
	sqlStatement := `
DELETE FROM people
WHERE id = $1;`
	_, err := Db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("People Deleted", id)

}

//Update People on ID
func updateName(name string, id int) {
	sqlStatement := `
UPDATE people
SET Name = $2
WHERE id = $1;`
	_, err := Db.Exec(sqlStatement, id, name)
	if err != nil {
		panic(err)
	}
	fmt.Println("People Name Updated", id)

}

//Drop Entire Table
func dropTable() {
	sqlStatement := `
	DROP TABLE IF EXISTS people;`
	_, err := Db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("People Table Dropped")

}

// func fetchAndAddAll() {
// 	resp, err := http.Get("https://swapi.dev/api/people/")
// 	fmt.Println("Fetching from https://swapi.dev/api/people/")
// 	if err != nil {
// 		log.Fatal("Error")
// 	}
// 	type Response struct {
// 		count interface{} `json:"results"`
// 	}
// 	r := Response{}

// 	err = json.NewDecoder(resp.Body).Decode(&r)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Print(r)
// 	// for  := range r.results {

// 	// }
// }
