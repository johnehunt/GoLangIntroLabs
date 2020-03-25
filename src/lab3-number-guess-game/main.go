package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game")

	// Initialise the random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialise the number to be guessed
	numberToGuess := random.Intn(10)

	// Initialise the number of tries the player has made
	countNumberOfTries := 1

	// Obtain their initial guess
	var guess int
	fmt.Print("Please guess a number between 1 and 10: ")
	fmt.Scanln(&guess)

	for {
		if numberToGuess != guess {
			fmt.Println("Sorry wrong number")
			if countNumberOfTries == 4 {
				break
			} else if guess < numberToGuess {
				fmt.Println("Your guess was lower than the number")
			} else {
				fmt.Println("Your guess was higher than the number")
			}
			fmt.Print("Please guess again (1 to 10): ")
			fmt.Scanln(&guess)
			countNumberOfTries++
		} else {
			break
		}
	}

	if numberToGuess == guess {
		fmt.Println("Well done you won!")
		fmt.Printf("You took %d goes to complete the game\n", countNumberOfTries)
	} else {
		fmt.Println("Sorry - you loose")
		fmt.Printf("The number you needed to guess was %d\n", numberToGuess)
	}

	fmt.Println("Game Over")
}
