package main

import (
	"flag"
	"log"
	"time"
)

// expected date format (DD/MM/YYYY)
var expectedDateFormat string = "02/01/2006"

// parseDate validate and parses a given date string.
func parseDate(date string) time.Time {
	parsedDate, err := time.Parse(expectedDateFormat, date)
	if err != nil || time.Now().After(parsedDate) {
		log.Fatal("Invalid date")
	}

	return parsedDate
}

// calcSleeps returns nights until the given date.
func calcSleeps(targetDate time.Time) (nights int) {
	return int(targetDate.Sub(time.Now()).Hours()) / 24
}

func main() {
	bday := flag.String("bday", "", "Your next birthday in DD/MM/YYYY format")
	flag.Parse()
	target := parseDate(*bday)
	log.Printf("You have %d sleeps until your birthday!", calcSleeps(target))
}
