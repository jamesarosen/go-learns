package main

import (
	"bufio"
	"fmt"
	"os"
)

func countWords(r Reader) map[string]int {
	wordCounts := map[string]int{}
	// adapted from https://golang.org/pkg/bufio/#example_Scanner_custom
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
        wordCounts[scanner.Text()]++
	}
	return wordCounts
}

func main() {
	for k, v := range countWords(os.Stdin) {
		fmt.Printf("%s: %d\n", k, v)
	}
}