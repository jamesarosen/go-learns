package main

import (
	"io"
	"os"
)

func main() {
	path := os.Args[1]
	println(path)
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var out io.Writer
	if (os.Args[2] == "--stderr") {
		out = os.Stderr
	} else {
		out = os.Stdout
	}

	io.Copy(out, file)
}