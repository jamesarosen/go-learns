package main

func main() {
	c := make(chan int)
	go compute(c)
	println(<-c)
}

func compute(c chan int) {
	c <- 123
}