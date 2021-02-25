package models

import (
	"database/sql"
	"log"
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

//CreateTable according to Struct and Swapi scheema
func (p People) CreateTable(Db *sql.DB) {
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
	log.Println("Table Created")
}

//Insert People to DB
func (p People) Insert(id int, Db *sql.DB) {
	sqlStatement := `
	INSERT INTO people (id,name , height , mass , haircolor , skincolor , eyecolor , birthyear , gender , homeworld , url , created , edited )
	VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	RETURNING id`
	err := Db.QueryRow(sqlStatement, id, p.Name, p.Height, p.Mass, p.HairColor, p.SkinColor, p.EyeColor, p.BirthYear, p.Gender, p.Homeworld, p.URL, p.Created, p.Edited).Scan(&p.ID)
	if err != nil {
		panic(err)
	}
	log.Println("People Inserted", p.ID)

}

//Delete People on ID
func (p People) Delete(id int, Db *sql.DB) {
	sqlStatement := `
DELETE FROM people
WHERE id = $1;`
	_, err := Db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	log.Println("People Deleted", id)

}

//Update People on ID
func (p People) Update(name string, id int, Db *sql.DB) {
	sqlStatement := `
UPDATE people
SET Name = $2
WHERE id = $1;`
	_, err := Db.Exec(sqlStatement, id, name)
	if err != nil {
		panic(err)
	}
	log.Println("People Name Updated", id)

}

//Drop Entire Table
func (p People) Drop(Db *sql.DB) {
	sqlStatement := `
	DROP TABLE IF EXISTS people;`
	_, err := Db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	log.Println("People Table Dropped")

}
