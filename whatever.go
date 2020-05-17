package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gorilla/mux"
)

type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string `gorm:"unique"`
	UserID uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})

	// Create
	db.Create(&User{Name: "Camille"})
	var user User
	db.First(&user, 1)
	db.Create(&CreditCard{Number: "4444333322221111", UserID: user.ID})

	// db.First(&expense, "code = ?", "L1212") // find expense with code l1212

	// // Update - update expense's price to 2000
	// db.Model(&expense).Update("Price", 2000)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	fmt.Println("Now listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(writer http.ResponseWriter, request *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Fetch Camille
	var user User
	db.First(&user, 1)

	var credit_cards []CreditCard
	db.Model(&user).Related(&credit_cards)
	json.NewEncoder(writer).Encode(credit_cards)
}
