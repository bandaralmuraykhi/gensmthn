//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the "report.json" file for reading.
	jsonFile, err := os.Open("report.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Defer the closing of our jsonFile to ensure it's closed later.
	defer jsonFile.Close()

	// Read the entire JSON file into a byte slice.
	byteValue, _ := io.ReadAll(jsonFile)

	// Create a map to hold the JSON data.
	var result map[string]interface{}

	// Unmarshal the JSON data into the result map.
	json.Unmarshal(byteValue, &result)

	// Extract the "Genre" and "Movie" parts from the result map.
	genres := result["Genre"].(map[string]interface{})
	movies := result["Movie"].(map[string]interface{})

	// Print genre earnings.
	fmt.Println("Genre Earnings:")
	for k, v := range genres {
		fmt.Printf("%-30s %v\n", k, v)
	}

	// Print movie earnings.
	fmt.Println("\nMovie Earnings:")
	for k, v := range movies {
		fmt.Printf("%-30s %v\n", k, v)
	}
}
