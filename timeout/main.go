package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tomatoTimer() string {
	delay := 1000 + rand.Intn(1000)
	time.Sleep(time.Duration(delay) * time.Millisecond)
    return fmt.Sprintf("ding! (waited %dms)", delay)
}

func main() {
	rand.Seed(2)

	c := make(chan string)
	go func() { c <- tomatoTimer() }()

	select {
	case x := <- c:
		println("Done", x)
	case <- time.After(1500 * time.Millisecond):
		println("Timeout")
	}
}