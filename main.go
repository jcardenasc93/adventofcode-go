package main

import (
	"fmt"

	"github.com/jcardenasc93/adventofcode-go/day1"
	"github.com/jcardenasc93/adventofcode-go/day2"
)

var option string

func main() {
	daySelection := getDayInput()
	switch daySelection {
	case "1":
		runDay1()
	case "2":
		runDay2()
	}
}
func getDayInput() string {
	fmt.Printf("Input day to solve: ")
	fmt.Scanf("%s", &option)
	fmt.Printf("Running solution for day %s\n", option)
	return option
}

func runDay1() {
	day1.RunDay1()
}

func runDay2() {
	day2.RunDay2()
}
