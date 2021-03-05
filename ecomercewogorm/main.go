package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rajaram"
	password = "rajaram"
	dbname   = "p1"
)

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
	log.Println("DB connected")
	// CreateDB(db)
	// Insert(Category{Name: "ROOT A"}, db)
	// Insert(Category{Name: "Child A1", Category_id: 1}, db)
	// Insert(Category{Name: "Child A2", Category_id: 1}, db)
	// Insert(Category{Name: "gChild A1A1", Category_id: 2}, db)
	// Insert(Category{Name: "gChild A1A2", Category_id: 2}, db)
	// Insert(Category{Name: "gChild A2A1", Category_id: 3}, db)
	// InsertProduct(Product{Name: "P1", Category_id: 3}, db)
	// InsertProduct(Product{Name: "P2", Category_id: 2}, db)

	getCategryWithID(1, db)

	// getProductByKeyProduct(1, db)
}

type Category struct {
	ID          int
	Name        string
	Category_id int64
	children    []Category
	Products    []Product
}

type Product struct {
	ID          int
	Name        string
	Category_id int64
	Variants    []Variant
}

type Variant struct {
	ID         int
	Name       string
	Product_id int64
}

func CreateDB(DB *sql.DB) {
	_, err := DB.Query(`
	CREATE TABLE categories (
		id        serial primary key,
		name      character varying NOT NULL,
		parent_id bigint 
	  );
	  CREATE INDEX ON categories(parent_id);
	  
	  CREATE TABLE products (
		id          serial primary key,
		name        character varying NOT NULL,
		category_id bigint NOT NULL
	  );
	  CREATE INDEX ON products(category_id);
	  
	  ALTER TABLE ONLY categories
		  ADD CONSTRAINT category_parent_fk FOREIGN KEY (parent_id) REFERENCES categories(id);
	  ALTER TABLE ONLY products
		  ADD CONSTRAINT product_category_fk FOREIGN KEY (category_id) REFERENCES categories(id);
	  `)
	if err != nil {
		panic(err)
	}
	log.Println("Table Created")
}

func Insert(cat Category, Db *sql.DB) int {
	if cat.Category_id == 0 {
		sqlStatement := `INSERT INTO categories(name) VALUES ($1)RETURNING id;`
		err := Db.QueryRow(sqlStatement, cat.Name).Scan(&cat.ID)
		if err != nil {
			panic(err)
		}
		log.Println("People Inserted", cat.ID)

		return cat.ID
	}
	sqlStatement := `INSERT INTO categories(name,parent_id) VALUES ($1,$2)RETURNING id;`
	err := Db.QueryRow(sqlStatement, cat.Name, cat.Category_id).Scan(&cat.ID)
	if err != nil {
		panic(err)
	}
	log.Println("category Inserted", cat.ID)
	return cat.ID

}

func InsertProduct(pro Product, Db *sql.DB) int {

	sqlStatement := `INSERT INTO products (name,category_id) VALUES ($1,$2) RETURNING id;`
	err := Db.QueryRow(sqlStatement, pro.Name, pro.Category_id).Scan(&pro.ID)
	if err != nil {
		panic(err)
	}
	log.Println("product Inserted", pro.ID)
	return pro.ID

}

func getProductByKeyProduct(cid int, Db *sql.DB) []Product {

	sqlStatement := `SELECT * from  products WHERE category_id=$1 ;`
	var res interface{}
	rows, err := Db.Query(sqlStatement, cid)
	if err != nil {
		panic(err)
	}
	products := []Product{}
	for rows.Next() {
		var id int
		var name string
		var category_id int
		err = rows.Scan(&id, &name, &category_id)
		if err != nil {
			// handle this error
			panic(err)
		}
		products = append(products, Product{ID: id, Name: name, Category_id: int64(category_id)})
		log.Println("product Inserted", res)
	}
	log.Println(products)
	return products
}

func getCategryWithID(id int, Db *sql.DB) {
	sqlStatement := `WITH RECURSIVE categories_with_roots AS (
		SELECT id, parent_id, name, name as root_name
		FROM categories
		WHERE parent_id IS NULL AND id=$1
	  
		UNION ALL
	  
		SELECT cat.id, cat.parent_id, cat.name, cwr.root_name
		FROM categories cat, categories_with_roots cwr
		WHERE cat.parent_id = cwr.id
	  )
	  SELECT id, name, parent_id FROM categories_with_roots;`

	rows, err := Db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id int
		var name string
		var category_id interface{}
		err = rows.Scan(&id, &name, &category_id)
		if err != nil {
			// handle this error
			panic(err)
		}

		// log.Println(id, name, category_id)
		if category_id == nil {
			categories = append(categories, Category{
				ID:          id,
				Name:        name,
				Category_id: 0,
			})
		} else {
			categories = append(categories, Category{
				ID:          id,
				Name:        name,
				Category_id: category_id.(int64),
			})
		}
	}
	// log.Println(categories)
	getStructuredCategory(categories, Db)
	return
}

func getStructuredCategory(categories []Category, Db *sql.DB) {
	for i := len(categories) - 1; i > 0; i-- {
		var x Category
		x, categories = categories[len(categories)-1], categories[:len(categories)-1]
		x.Products = getProductByKeyProduct(x.ID, Db)
		log.Println(x.Category_id, x.ID, x)
		for j := len(categories) - 1; j > -1; j-- {
			if categories[j].ID == int(x.Category_id) {
				categories[j].children = append(categories[j].children, x)
				break
			}
		}
		if x.Category_id == 0 {
			break
		}
	}
	log.Println(categories[0])

}
