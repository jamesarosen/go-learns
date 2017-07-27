package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := map[string]int{}

	// adapted from https://golang.org/pkg/bufio/#example_Scanner_custom
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
        wordCounts[scanner.Text()]++
	}

	for k, v := range wordCounts {
		fmt.Printf("%s: %d\n", k, v)
	}
}