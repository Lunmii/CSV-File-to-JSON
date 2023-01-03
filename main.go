package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	csvFile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	// Parse the CSV file
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the CSV records to JSON
	jsonData, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the JSON data to a file
	jsonFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}
