package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
type Player struct {
  gorm.Model
	Name string `gorm:"unique;not null"`
	Race string
	Class string
	Items []Item `gorm:"foreignkey:UserRefer"`
}

type Item struct {
  gorm.Model
	Name string `gorm:"unique;not null"`
	Description string `gorm:"unique;not null"`
	Rarity string
	Attack string
}

func main() {
  db, err := gorm.Open("postgres", "host=localhost port=35432 user=user dbname=db password=pass sslmode=disable")
  if err != nil {
    fmt.Println(err)
    panic("failed to connect database")
  }

	fmt.Println("connected to db")

  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Player{}, &Item{})

  // Create
	 var players []Player = []Player{
    Player{Name: "Haeli", Race: "Centaur", Class: "Cleric"},
    Player{Name: "Andrew", Race: "Elf", Class: "Wizard/Barbarian"},
	}

	for _, player := range players {
	  db.Create(&player)
		fmt.Println("Created player: ", player.Name)
	}

// Read
// var product Product
// db.First(&product, 1) // find product with id 1
// db.First(&product, "code = ?", "L1212") // find product with code l1212
// 
// fmt.Println(product)
//
// // Update - update product's price to 2000
// db.Model(&product).Update("Price", 2000)
//
// // Delete - delete product
// db.Delete(&product)
}
