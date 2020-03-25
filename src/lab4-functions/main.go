package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getUserInput(guess *int, prompt string) {
	fmt.Print(prompt)
	fmt.Scanln(guess)
}

func playGame() {
	// Initialise the random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialise the number to be guessed
	numberToGuess := random.Intn(10)

	// Initialise the number of tries the player has made
	countNumberOfTries := 1

	// Obtain their initial guess
	var guess int
	getUserInput(&guess, "Please guess a number between 1 and 10: ")

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
			getUserInput(&guess, "Please guess again (1 to 10): ")
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
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game")

	playGame()

	fmt.Println("Game Over")
}
