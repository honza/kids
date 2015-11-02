package main

import (
	"bufio"
	"fmt"
	"github.com/honza/kids/lib"
	"os"
	"strconv"
	"strings"
)

func GetNumberFromUser() int {

	for {

		fmt.Print("Your guess: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		i, err := strconv.Atoi(strings.Trim(text, "\n"))

		if err != nil {
			fmt.Println("Not a number.")
			continue
		}

		return i

	}
}

func main() {
	secret := lib.RandomInt()

	fmt.Println("Guess a number between 0 and 100.")

	for {
		n := GetNumberFromUser()

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
