package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func readCSV(f string) {
	csvFile, _ := os.Open(f)
	reader := csv.NewReader(csvFile)

	var people []Person
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}

	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}

func TestReadCSV(t *testing.T) {
	fmt.Println("\nTestReadCSV")
	readCSV("./data.csv")
}
