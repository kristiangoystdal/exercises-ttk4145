// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

var i = 0

func server(operation chan string) {
	for {
		select {
		case op := <-operation:
			if op == "increment" {
				i++
			} else if op == "decrement" {
				i--
			}
		default:
		}
	}
}

func incrementing(operation chan string, finish chan bool) {
	//TODO: increment i 1000000 times
	for n := 0; n < 1042; n++ {
		operation <- "increment"
	}
	finish <- true

}

func decrementing(operation chan string, finish chan bool) {
	//TODO: decrement i 1000000 times
	for n := 0; n < 1000; n++ {
		operation <- "decrement"
	}
	finish <- true
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	operation := make(chan string)
	finish := make(chan bool)

	go server(operation)

	// TODO: Spawn both functions as goroutines
	go incrementing(operation, finish)
	go decrementing(operation, finish)

	<-finish
	<-finish
	fmt.Println("The magic number is:", i)
}
