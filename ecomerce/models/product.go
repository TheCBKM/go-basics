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

type Product struct {
	gorm.Model
	Name        string     `json:"Name" gorm:"not null;default:null"`
	Description string     `json:"Description"`
	ImageURL    string     `json:"ImageURL"`
	CategoryID  int        `json:"CategoryID"`
	Variants    []*Variant `json:"Variants" gorm:"not null;default:null"`
}

func (p Product) Insert(db *gorm.DB) {
	result := db.Create(&p)
	if result.Error != nil {
		log.Panic("Inser error")
	}

}

func (p Product) Update(db *gorm.DB) {
	db.Save(&p)
}

func (p Product) Delete(id int, db *gorm.DB) {
	db.Delete(&Product{}, id)
}

func (p *Product) Get(id int, db *gorm.DB) {
	db.First(&p, id)

}

func (c *Product) GetAll(db *gorm.DB) []Product {
	cs := []Product{}
	db.Find(&cs)
	return cs
}

func (p *Product) RegisterHandlers(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/get/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		getProduct(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/getall/product", func(w http.ResponseWriter, r *http.Request) {
		getAllProducts(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/create/product", func(w http.ResponseWriter, r *http.Request) {
		createProduct(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/delete/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteProduct(w, r, db)
	}).Methods("DELETE")

}

func getProduct(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Product{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Get(y, db)
	json.NewEncoder(w).Encode(&c)
}

func deleteProduct(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Product{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Delete(y, db)
	json.NewEncoder(w).Encode(&c)
}

func createProduct(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	p := Product{}
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

func getAllProducts(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Product{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	cs := c.GetAll(db)
	json.NewEncoder(w).Encode(&cs)
}
