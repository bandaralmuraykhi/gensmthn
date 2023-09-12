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
	totalm := map[string]int{}
	totalg := map[string]int{}

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
		// year, _ := strconv.Atoi(fields[2])
		earnings, _ := strconv.Atoi(fields[3])

		//fmt.Println(movie)

		totalm[movie] += earnings
		totalg[genre] += earnings

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	// for k, v := range totalm {
	// 	fmt.Println("Movie:", k, "eranings:", v)
	// }
	// fmt.Println("########################################")

	// for k, v := range totalg {
	// 	fmt.Println("Genre:", k, "eranings:", v)
	// }

	var report struct {
		Movie map[string]int
		Genre map[string]int
	}
	report.Movie = totalm
	report.Genre = totalg
	text, _ := json.MarshalIndent(report, "", " ")
	// text, _ := json.Marshal(report)
	fmt.Println(string(text))

	_ = ioutil.WriteFile("report.json", text, 0644)
}
