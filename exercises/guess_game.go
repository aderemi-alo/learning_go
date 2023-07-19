package exercises

import (
	"fmt"
	"math/rand"
)

// A guessing game, where the user has 10 guesses to guess the number
func GuessGame() {
	target := rand.Intn(101)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// Loop to allow the user guess 10 times
	for i := 1; i < 11; i++ {
		fmt.Printf("Guess #%d: ", i)
		var guess int
		fmt.Scanf("%d", &guess)
		if guess == target {
			fmt.Println("You got it!")
			break
		} else if guess < target {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
		if i == 10 {
			fmt.Printf("Game over!")
		}
	}
}
