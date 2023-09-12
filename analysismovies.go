//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// YE represents the year and earnings for a movie.
type YE struct {
	year     int
	earnings int
}

// Entry represents an entry for a movie in the dataset.
type Entry struct {
	movie    string
	year     int
	earnings int
}

func main() {
	entries := []Entry{}
	allmovies := map[string]*YE{}
	allgenres := map[string]int{}

	// Open the "moviesdata.csv" file for reading.
	f, err := os.Open("moviesdata.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	// Skip the first line (header) if it exists.
	if scanner.Scan() {
		// Skip the header line without printing it
	}

	// Read each subsequent line of the file.
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ", ")
		genre := fields[0]
		movie := fields[1]
		year, _ := strconv.Atoi(fields[2])
		earnings, _ := strconv.Atoi(fields[3])

		// Update the earnings for the genre.
		allgenres[genre] += earnings

		// Update the year and earnings for the movie.
		value, ok := allmovies[movie]
		if !ok {
			value = &YE{}
			allmovies[movie] = value
		}
		value.year = year
		value.earnings = earnings
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	// Process the data and create entries.
	for key, val := range allmovies {
		entries = append(entries, Entry{key, val.year, val.earnings})
	}

	// Print the total earnings for each genre.
	for genre, earnings := range allgenres {
		fmt.Printf("Genre: %s, Total Earnings: %d\n", genre, earnings)
	}
}
