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
	Name        string `gorm:"unique"`
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
	err = db.Create(&User{Name: "Camille"}).Error
	if err != nil {
		fmt.Printf("couldn't create user: %v\n", err)
	}
	var user User
	db.Where(&User{Name: "Camille"}).First(&user)
	db.Create(&CreditCard{Number: "4444333322221111", UserID: user.ID})
	db.Create(&CreditCard{Number: "1010292938384747", UserID: user.ID})
	db.Create(&CreditCard{Number: "1234567812345678", UserID: user.ID})

	// db.First(&expense, "code = ?", "L1212") // find expense with code l1212

	// // Update - update expense's price to 2000
	// db.Model(&expense).Update("Price", 2000)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/users", ShowAllUsers)
	router.HandleFunc("/users/{name}", ShowUser)
	// router.HandleFunc("/users", Method: "post", CreateUser)
	fmt.Println("Now listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// func CreateUser(Name string) {
// 	db.FirstOrCreate(&user, User{Name: "non_existing"})

// 	return 201 CREATED
// }

func ShowUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	params := mux.Vars(r)
	name := params["name"]

	var user User
	db.Where(&User{Name: name}).First(&user)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users User
	db.Find(&users)

	if err != nil {
		log.Fatalf("Unable to get users. %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

func Index(writer http.ResponseWriter, request *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Fetch Camille
	var user User
	db.Preload("CreditCards").Where(&User{Name: "Camille"}).First(&user)
	json.NewEncoder(writer).Encode(user)
	
	// var credit_cards []CreditCard
	// db.Model(&user).Related(&credit_cards)
	// json.NewEncoder(writer).Encode(credit_cards)
}
