package main

import (
	"fmt"
	"time"
)

// LOOP_TOTAL used for number of lops run by worker
const LOOP_TOTAL = 20

// Message structure used to send messages between
// the producer and worker functions
type Message struct {
	ID  int
	Msg string
}

func worker(requests <-chan Message) {
	for message := range requests {
		fmt.Print(message.Msg)
	}
}

func producer(msg string, requests chan<- Message) {
	for i := 0; i < LOOP_TOTAL; i++ {
		requests <- Message{i, msg}
		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting")

	fmt.Println("Setting up requests channel")
	request := make(chan Message, 4)

	fmt.Println("Launching worker Goroutine")
	go worker(request)

	fmt.Println("Launching four producer Goroutines")
	go producer("A", request)
	go producer("B", request)
	go producer("C", request)
	go producer("D", request)

	fmt.Println("Wait for a second")
	time.Sleep(time.Second)

	close(request)

	fmt.Println("\nDone")
}
