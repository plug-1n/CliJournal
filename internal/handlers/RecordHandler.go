package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/plug-1n/cli-journal/internal/models"
)

func GetAllRecords() ([]byte, error) {
	content, err := ioutil.ReadFile("temp_db.json")
	if err != nil {
		log.Fatal("Error during opening file: ", err)
	}
	var payLoad []models.Record
	err = json.Unmarshal(content, &payLoad)
	if err != nil {
		log.Fatal("Error during Unmarshall: ", err)
	}
	return content, nil

}
func AddNewRecord() {
	fmt.Print("Input title: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if title == "" {
		log.Fatal("Title can't be empty")
	}
	fmt.Println("Input content: ")
	body, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now().Format("2006-01-02")

	content, err := GetAllRecords()
	if err != nil {
		log.Fatal("Error during Get All records: ", err)
	}
	id := uuid.New()
	newRecord = models.Record(id, title, body, currentTime)
	content = append(content, []string{title, body, currentTime})

}
