package main

import (
	"fmt"

	"github.com/jcardenasc93/adventofcode-go/day1"
	"github.com/jcardenasc93/adventofcode-go/day2"
	"github.com/jcardenasc93/adventofcode-go/day3"
)

var option string

func main() {
	daySelection := getDayInput()
	switch daySelection {
	case "1":
		runDay1()
	case "2":
		runDay2()
	case "3":
		runDay3()
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

func runDay3() {
	day3.RunDay3()
}
