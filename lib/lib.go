package lib

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func Panic(err error) {
	fmt.Println("*******************************************************************************")
	fmt.Println("Something went wrong, and the program died.\n")
	fmt.Println("Here is the problem the program had:\n")
	fmt.Println(err.Error())
	fmt.Println("*******************************************************************************")
	os.Exit(1)
}

func Days() {
	fmt.Println("days")
}

func DurationToDays(d time.Duration) int {
	hours := int(d.Hours())
	return int(hours / 24)
}

func DaysUntil(t time.Time) int {
	now := time.Now()
	delta := t.Sub(now)
	return DurationToDays(delta)
}

func RandomInt(ceil int) int {
	i := big.NewInt(int64(ceil))
	value, err := rand.Int(rand.Reader, i)

	if err != nil {
		Panic(err)
	}

	return int(value.Int64())

}

func GetNumberFromUser(prompt string) int {

	for {

		fmt.Print(prompt)
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
