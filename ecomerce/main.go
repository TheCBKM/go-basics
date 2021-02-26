package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rajaram"
	password = "rajaram"
	dbname   = "practice_db"
)

var (
	categories = []Category{
		{Name: "Men"},
		{Name: "Women"},
	}

	products = []Product{
		{Name: "P1", CategoryID: 1},
		{Name: "P2", CategoryID: 2},
	}

	variant = []Variant{
		{Name: "v1", ProductID: 1},
		{Name: "v22", ProductID: 2},
	}
)

var db *gorm.DB
var err error

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("DB connected")
	defer db.Close()

	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Variant{})

	// for index := range categories {
	// 	categories[index].Insert(db)
	// }

	// for index := range products {
	// 	products[index].Insert(db)
	// }

	// for index := range variant {
	// 	variant[index].Insert(db)
	// }

	router := mux.NewRouter()
	router.HandleFunc("/get/category/{id}", getCategory).Methods("GET")
	router.HandleFunc("/get/product/{id}", getProduct).Methods("GET")
	router.HandleFunc("/get/variant/{id}", getVariant).Methods("GET")
	router.HandleFunc("/create/category", create).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	c := Category{}
	log.Println("WOW")
	if err := decoder.Decode(&c); err != nil {
		log.Panic(err)
		return
	}
	log.Println(c)
	defer r.Body.Close()

	c.Insert(db)
	json.NewEncoder(w).Encode("Done")

}

func getCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	c := Category{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Get(y, db)
	json.NewEncoder(w).Encode(&c)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	c := Product{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Get(y, db)
	json.NewEncoder(w).Encode(&c)
}

func getVariant(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	c := Variant{}
	y, e := strconv.Atoi(params["id"])
	if e == nil {
		fmt.Printf("%T \n %v", y, y)
	}
	c.Get(y, db)
	json.NewEncoder(w).Encode(&c)
}
