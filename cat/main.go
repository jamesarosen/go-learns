package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func main() {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	path := os.Args[1]
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	f.Write(dat)
}