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
	christmas := time.Date(thisYear, time.December, 25, 0, 0, 0, 0, location)
	daysLeft := lib.DaysUntil(christmas)

	if daysLeft == 1 {
		fmt.Printf("%d day left until Christmas\n", daysLeft)
	} else {
		fmt.Printf("%d days left until Christmas\n", daysLeft)
	}

}
