package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Variant struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"Name" gorm:"not null;default:null"`
	MRP           int64  `json:"MRP"  gorm:"not null;default:null"`
	DiscountPrice int64  `json:"DiscountPrice"`
	Size          uint   `json:"Size"`
	Colour        string `json:"Colour"`
	ProductID     int    `json:"ProductID"`
}

func (v Variant) Insert(db *gorm.DB) {
	result := db.Create(&v)
	if result.Error != nil {
		log.Panic("Insert error")
	}

}

func (v Variant) Update(db *gorm.DB) {
	db.Save(&v)
}

func (v Variant) Delete(id int, db *gorm.DB) {
	db.Delete(&Variant{}, id)
}

func (v *Variant) Get(id int, db *gorm.DB) {
	db.First(&v, id)
}

func (c *Variant) GetAll(db *gorm.DB) []Variant {
	cs := []Variant{}
	db.Find(&cs)
	return cs
}

func (p *Variant) RegisterHandlers(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/get/variant/{id}", func(w http.ResponseWriter, r *http.Request) {
		getVariant(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/getall/variant", func(w http.ResponseWriter, r *http.Request) {
		getAllVariants(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/create/variant", func(w http.ResponseWriter, r *http.Request) {
		createVariant(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/delete/variant/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteVariant(w, r, db)
	}).Methods("DELETE")
}

func createVariant(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	p := Variant{}
	log.Println("WOW")
	if err := decoder.Decode(&p); err != nil {
		log.Panic(err)
		return
	}
	log.Println(p)
	defer r.Body.Close()

	p.Insert(db)
	json.NewEncoder(w).Encode("Done")

}

func getVariant(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Variant{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Get(y, db)
	json.NewEncoder(w).Encode(&c)
}

func deleteVariant(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Variant{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Delete(y, db)
	json.NewEncoder(w).Encode(&c)
}

func getAllVariants(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Variant{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	cs := c.GetAll(db)
	json.NewEncoder(w).Encode(&cs)
}
