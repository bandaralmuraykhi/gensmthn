//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	jsonFile, err := os.Open("report.json")

	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	s := result["Genre"].(map[string]interface{})

	for k, v := range s {
		fmt.Printf("%-30s %v\n", k, v)
	}
	// ################## different way to do it #####################
	// // fmt.Println(report)
	// for k, v := range report.Movie {
	// 	fmt.Println(k, v)
	// }
	// for k, v := range report.Genre {
	// 	fmt.Println(k, v)
	// }

}
