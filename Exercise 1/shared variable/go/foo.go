// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"sync"
	"time"
)

var i = 0
var lock sync.Mutex 

func incrementing() {
	//TODO: increment i 1000000 times
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for n := 0; n < 1000042; n++ {
			lock.Lock()
			i += 1
			lock.Unlock()
		}
	}()
}

	

func decrementing() {
	//TODO: decrement i 1000000 times
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for n := 0; n < 1000000; n++ {
			lock.Lock()
			i -= 1
			lock.Unlock()
		}
	}()
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	incrementing()
	decrementing()

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
