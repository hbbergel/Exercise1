// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    for t := 0; t < 10000000; t++ {
        i++
    }
    //TODO: increment i 1000000 times
}

func decrementing() {
    //TODO: decrement i 1000000 times
    for t := 0; t < 10000000; t++ {
        i--
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    
    go incrementing()
    go decrementing()
    // I guess this is a hint to what GOMAXPROCS does...
	                                    // Try doing the exercise both with and without it!

    // TODO: Spawn both functions as goroutines
	
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println("The magic number is:", i)
}
