package main

import "time"

func main() {
	foo("a")
	go foo("b")
	time.Sleep(time.Second)
}

func foo(s string) {
	println(s)
}