//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type YE struct {
	year     int
	earnings int
}
type Entry struct {
	movie    string
	year     int
	earnings int
}

func main() {
	entries := []Entry{}
	allmovies := map[string]*YE{}
	allgenres := map[string]int{}

	f, err := os.Open("moviesdata.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ", ")
		genre := fields[0]
		movie := fields[1]
		year, _ := strconv.Atoi(fields[2])
		earnings, _ := strconv.Atoi(fields[3])

		fmt.Println(genre)

		allgenres[genre] += earnings

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
	// data processing
	for key, val := range allmovies {
		entries = append(entries, Entry{key, val.year, val.earnings})
	}
	// fmt.Println(entries)
	fmt.Println(allgenres)
}
