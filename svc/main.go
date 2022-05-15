package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Town struct {
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Name string
}

func main() {

	dsn := "host=localhost user= password= dbname= port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Town{})
	// db.Migrator().CreateTable(&Town{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "add")
		town := Town{Name: "%q"}
		db.Create(&town)
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "all")
		town := Town{}
		db.First(&town)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
