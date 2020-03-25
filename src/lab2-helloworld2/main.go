package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello Johns' Bookshop")

	var user string

	fmt.Print("Please enter your name: ")
	fmt.Scanln(&user)

	now := time.Now()

	fmt.Printf("Welcome %s today is %v\n", user, now)

}
