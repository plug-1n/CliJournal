package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/plug-1n/cli-journal/internal/models"
)

func GetAllData() {
	content, err := ioutil.ReadFile("temp_db.json")
	if err != nil {
		log.Fatal("Error during opening file: ", err)
	}
	var payLoad []models.Record
	err = json.Unmarshal(content, &payLoad)
	if err != nil {
		log.Fatal("Error during Unmarshall: ", err)
	}
	log.Println(payLoad)

}
