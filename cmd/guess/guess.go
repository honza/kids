package main

import (
	"fmt"
	"github.com/honza/kids/lib"
)

func main() {
	secret := lib.RandomInt(100)

	fmt.Println("Guess a number between 0 and 100.")

	for {
		n := lib.GetNumberFromUser("Your guess: ")

		if n == secret {
			fmt.Println("You win!")
			return
		}

		if n < secret {
			fmt.Println("Too small!")
			continue
		}

		if n > secret {
			fmt.Println("Too big!")
			continue
		}
	}

}
