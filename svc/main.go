package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Town struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	City string `gorm:"not null;unique" json:"city"`
}

func main() {

	dsn := "host=db-service user=postgres password=postgres_secret dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
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
		var town Town
		err := json.NewDecoder(r.Body).Decode(&town)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db.Create(&town)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "all")
		town := []Town{}
		db.Find(&town)

		output := make([]string, 0, len(town))
		for _, element := range town {
			output = append(output, element.City)
		}

		js, err := json.Marshal(output)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
