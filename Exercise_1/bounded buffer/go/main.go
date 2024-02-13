package main

import (
	"fmt"
	"time"
)

func producer(buffer chan int) {

	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		buffer <- i
	}

}

func consumer(buffer chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-buffer //TODO: get real value from buffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	buffer := make(chan int, 5)

	go consumer(buffer)
	go producer(buffer)

	select {}
}
