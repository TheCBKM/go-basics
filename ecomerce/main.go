package main

import (
	"fmt"
	"log"
	"net/http"

	models "ecomerce/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Postgres data
const (
	host     = "localhost"
	port     = 5432
	user     = "rajaram"
	password = "rajaram"
	dbname   = "practice_db"
)

//Initial data
var (
	categories = []models.Category{
		{Name: "Men"},
		{Name: "Women"},
	}

	products = []models.Product{
		{Name: "P1", CategoryID: 1},
		{Name: "P2", CategoryID: 2},
		{Name: "P3", CategoryID: 1},
		{Name: "P4", CategoryID: 2},
	}

	variant = []models.Variant{
		{Name: "v1", ProductID: 1, MRP: 300},
		{Name: "v2", ProductID: 2, MRP: 200},
	}
)

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("DB connected")
	defer db.Close()
	//Migrate DB
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Variant{})
	//Preload all pointers
	db = db.Set("gorm:auto_preload", true)
	//inset data
	for index := range categories {
		categories[index].Insert(db)
	}

	for index := range products {
		products[index].Insert(db)
	}

	for index := range variant {
		variant[index].Insert(db)
	}

	router := mux.NewRouter()
	c := models.Category{}
	p := models.Product{}
	v := models.Variant{}

	c.RegisterHandlers(router, db)
	p.RegisterHandlers(router, db)
	v.RegisterHandlers(router, db)

	handler := cors.Default().Handler(router)
	log.Println("Hosted on localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}
