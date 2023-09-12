//go:build ignore

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalm := map[string]int{} // Map to store total earnings per movie
	totalg := map[string]int{} // Map to store total earnings per genre

	// Open the "moviesdata.csv" file for reading.
	f, err := os.Open("moviesdata.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	// Read each line of the file.
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ", ")
		genre := fields[0]
		movie := fields[1]
		earnings, _ := strconv.Atoi(fields[3])

		// Update total earnings for each movie and genre.
		totalm[movie] += earnings
		totalg[genre] += earnings
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	// Create a report struct to store the maps.
	var report struct {
		Movie map[string]int
		Genre map[string]int
	}
	report.Movie = totalm
	report.Genre = totalg

	// Marshal the report struct to JSON with indentation.
	text, _ := json.MarshalIndent(report, "", " ")

	// Print the JSON report.
	fmt.Println(string(text))

	// Write the JSON report to a file named "report.json."
	_ = ioutil.WriteFile("report.json", text, 0644)
}
