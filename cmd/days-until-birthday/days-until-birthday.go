package main

import (
	"fmt"
	"github.com/honza/kids/lib"
	"time"
)

func main() {
	now := time.Now()

	location := now.Location()
	thisYear := now.Year()

	birthday := time.Date(thisYear, time.November, 13, 0, 0, 0, 0, location)
	born := time.Date(2009, time.November, 13, 0, 0, 0, 0, location)

	if !now.Before(birthday) {
		birthday = time.Date(thisYear+1, time.November, 13, 0, 0, 0, 0, location)
	}

	daysLeft := lib.DaysUntil(birthday)
	age := int(birthday.Sub(born).Hours() / 24 / 365)

	if daysLeft == 1 {
		fmt.Printf("%d day left until your birthday\n", daysLeft)
	} else {
		fmt.Printf("%d days left until your birthday\n", daysLeft)
	}

	fmt.Printf("You will be %d years old.\n", age)

}
