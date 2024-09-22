package db

import (
	"encoding/json"
	"log"
	"os"
)

func (db *Database) Seed() {
	file, err := os.ReadFile("./db/cars.json")
	if err != nil {
		log.Fatal(err)
	}

	var cars []CarCharacteristics
	err = json.Unmarshal(file, &cars)
	if err != nil {
		log.Fatal(err)
	}

	var chunks [][]CarCharacteristics

	for i := 0; i < len(cars); i += 100 {
		chunks = append(chunks, cars[i:min(i+100, len(cars))])
	}

	for _, chunk := range chunks {
		err = Db.CreateManyCars(chunk)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Database seeded")

	count, err := Db.CountCars()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(count)
}
