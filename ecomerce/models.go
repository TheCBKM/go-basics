package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Category ...
type Category struct {
	gorm.Model
	Name      string      `json:"Name"`
	ParentID  *uint       `json:"ParentID"`
	Categorys []*Category `json:"Categorys"  gorm:"foreignkey:ParentID" `
	Products  []*Product  `json:"Products"`
}

type Product struct {
	gorm.Model
	Name        string     `json:"Name"`
	Description string     `json:"Description"`
	ImageURL    string     `json:"ImageURL"`
	CategoryID  int        `json:"CategoryID"`
	Variants    []*Variant `json:"Variants"`
}

type Variant struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"Name"`
	MRP           int64  `json:"MRP"`
	DiscountPrice int64  `json:"DiscountPrice"`
	Size          uint   `json:"Size"`
	Colour        string `json:"Colour"`
	ProductID     int    `json:"ProductID"`
}

func (c Category) Insert(db *gorm.DB) {
	result := db.Create(&c)
	if result.Error != nil {
		log.Panic("Inser error")
	}
	log.Println("INserted")
}

func (c Category) Update(db *gorm.DB) {
	db.Save(&c)
}

func (c Category) Delete(id int, db *gorm.DB) {
	db.Delete(&Category{}, id)
}

func (c *Category) Get(id int, db *gorm.DB) {
	db.Preload("Categorys").Preload("Products").First(&c, id).Preload("Categorys").Preload("Products")
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
	db.Preload("Variants").First(&p, id)

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
