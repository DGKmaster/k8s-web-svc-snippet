package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Town struct {
	// ID is not used in return in /all endpoint
	ID uint `gorm:"primaryKey;autoIncrement" json:"-"`
	// City must be unique and not empty string
	City string `gorm:"not null;unique" json:"city"`
}

func main() {
	// DB connection
	// TODO: Change to ENV and Secret
	dsn := "host=db-service user=postgres password=postgres_secret dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
	// dsn := "host=localhost user=postgres password=postgres_secret dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Auto create table towns and update based on model
	db.AutoMigrate(&Town{})

	// Add new town ednpoint
	// If it is existed do not create duplicate
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var town Town
		err := json.NewDecoder(r.Body).Decode(&town)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db.Create(&town)

		w.Header().Set("Content-Type", "application/json")
		resp := "Completed"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	// List all cities endpoint
	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Limit output
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

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
