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

// Category ...
type Category struct {
	gorm.Model
	Name      string      `json:"Name" gorm:"not null;default:null"`
	ParentID  *uint       `json:"ParentID"`
	Categorys []*Category `json:"Categorys"  gorm:"foreignkey:ParentID `
	Products  []*Product  `json:"Products"`
}

func (c Category) Insert(db *gorm.DB) {
	result := db.Create(&c)
	if result.Error != nil {
		log.Panic("Insert error")
	}
	log.Println("INserted")
}

func (c Category) Update(db *gorm.DB) {
	db.Save(&c)
}

func (c Category) Delete(id int, db *gorm.DB) {
	db.First(&c, id)
	db.Delete(&c)
}

func (c *Category) Get(id int, db *gorm.DB) {
	db.First(&c, id)
}
func (c *Category) GetAll(db *gorm.DB) []Category {
	cs := []Category{}
	db.Find(&cs)
	return cs
}

func (c *Category) RegisterHandlers(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/get/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		getCategory(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/getall/category", func(w http.ResponseWriter, r *http.Request) {
		getAllCategory(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/create/category", func(w http.ResponseWriter, r *http.Request) {
		createCategory(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/delete/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteCategory(w, r, db)
	}).Methods("DELETE")
}

func getCategory(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Category{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Get(y, db)
	json.NewEncoder(w).Encode(&c)
}

func createCategory(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	c := Category{}
	if err := decoder.Decode(&c); err != nil {
		log.Panic(err)
		return
	}
	log.Println(c)
	defer r.Body.Close()

	c.Insert(db)
	json.NewEncoder(w).Encode("Done")

}

func deleteCategory(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Category{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	log.Println(y)
	c.Delete(y, db)
	json.NewEncoder(w).Encode(&c)
}

func getAllCategory(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	c := Category{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	cs := c.GetAll(db)
	json.NewEncoder(w).Encode(&cs)
}
