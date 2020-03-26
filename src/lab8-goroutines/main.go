package main

import (
	"fmt"
	"sync"
	"time"
)

// LoopTotal used for number of lops run by worker
const LoopTotal = 20

func worker(msg string, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we’re done.
	defer wg.Done()

	for i := 0; i < LoopTotal; i++ {
		fmt.Print(msg)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting")

	fmt.Println("Setting up the WaitGroup")
	var wg sync.WaitGroup
	wg.Add(4)

	fmt.Println("Launching four Goroutines")
	go worker("A", &wg)
	go worker("B", &wg)
	go worker("C", &wg)
	go worker("D", &wg)

	// Wait for goroutines to complete
	fmt.Println("Waiting for Goroutines to finish")
	wg.Wait()

	fmt.Println("\nDone")
}
