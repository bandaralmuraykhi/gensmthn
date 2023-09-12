//go:build ignore

package main

import (
	"fmt"
	"math/rand"
)

func getgenre() string {
	var genre = []string{"Comdey", "Drama", "Action"}
	return genre[rand.Intn(len(genre))]
}

func getmovie() string {
	var movie = []string{"The Shawshank Redemption", "American Psycho", "The Godfather", "The Dark Knight", "12 Angry Men", "Pulp Fiction", "Fight Club", "Inception", "The Matrix", "The Pianist"}
	return movie[rand.Intn(len(movie))]
}

func main() {
	fmt.Println("Genre, Movie, year, earnings")
	min := 1950
	max := 2023
	for i := 0; i < 20; i++ {
		fmt.Printf("%s, %s, %d, %d\n", getgenre(), getmovie(), rand.Intn(max-min)+min, rand.Intn(20000)+1000)
	}
}
