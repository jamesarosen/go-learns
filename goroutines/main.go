package main

import "time"

func main() {
	foo("a")
	go foo("b")
}

func foo(s string) {
	println(s)
}