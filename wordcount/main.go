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

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	    return bufio.ScanWords(data, atEOF)
	}

	scanner.Split(split)

	for scanner.Scan() {
        wordCounts[scanner.Text()]++
	}

	for k, v := range wordCounts {
		fmt.Printf("%s: %d\n", k, v)
	}
}