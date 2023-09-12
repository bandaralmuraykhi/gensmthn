//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getGenre() string {
	// Define a list of movie genres
	genres := []string{"Comedy", "Drama", "Action"}

	// Randomly select a genre
	return genres[rand.Intn(len(genres))]
}

func getMovie() string {
	// Define a list of popular movies
	movies := []string{
		"The Shawshank Redemption", "American Psycho", "The Godfather", "The Dark Knight",
		"12 Angry Men", "Pulp Fiction", "Fight Club", "Inception", "The Matrix", "The Pianist",
	}

	// Randomly select a movie
	return movies[rand.Intn(len(movies))]
}

func main() {
	// Seed the random number generator to ensure different results on each run
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Genre, Movie, Year, Earnings")

	minYear := 1950
	maxYear := 2023

	for i := 0; i < 20; i++ {
		// Generate random values for year and earnings
		year := rand.Intn(maxYear-minYear+1) + minYear
		earnings := rand.Intn(20000) + 1000

		// Call functions to get random genre and movie
		genre := getGenre()
		movie := getMovie()

		// Print the formatted data
		fmt.Printf("%s, %s, %d, %d\n", genre, movie, year, earnings)
	}
}
