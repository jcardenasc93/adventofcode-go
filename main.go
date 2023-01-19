package main

import (
	// "fmt"

	"github.com/jcardenasc93/adventofcode-go/day1"
	"github.com/jcardenasc93/adventofcode-go/day10"
	"github.com/jcardenasc93/adventofcode-go/day11"
	"github.com/jcardenasc93/adventofcode-go/day12"
	"github.com/jcardenasc93/adventofcode-go/day13"
	"github.com/jcardenasc93/adventofcode-go/day14"
	"github.com/jcardenasc93/adventofcode-go/day2"
	"github.com/jcardenasc93/adventofcode-go/day3"
	"github.com/jcardenasc93/adventofcode-go/day4"
	"github.com/jcardenasc93/adventofcode-go/day5"
	"github.com/jcardenasc93/adventofcode-go/day6"
	"github.com/jcardenasc93/adventofcode-go/day7"
	"github.com/jcardenasc93/adventofcode-go/day8"
	"github.com/jcardenasc93/adventofcode-go/day9"
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
	case "4":
		runDay4()
	case "5":
		runDay5()
	case "6":
		runDay6()
	case "7":
		runDay7()
	case "8":
		runDay8()
	case "9":
		runDay9()
	case "10":
		runDay10()
	case "11":
		runDay11()
	case "12":
		runDay12()
	case "13":
		runDay13()
	case "14":
		runDay14()
	}
}
func getDayInput() string {
	// fmt.Printf("Input day to solve: ")
	// fmt.Scanf("%s", &option)
	// fmt.Printf("Running solution for day %s\n", option)
	return "14"
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

func runDay4() {
	day4.RunDay4()
}

func runDay5() {
	day5.RunDay5()
}

func runDay6() {
	day6.RunDay6()
}

func runDay7() {
	day7.RunDay7()
}

func runDay8() {
	day8.RunDay8()
}

func runDay9() {
	day9.RunDay9()
}

func runDay10() {
	day10.RunDay10()
}

func runDay11() {
	day11.RunDay11()
}

func runDay12() {
	day12.RunDay12()
}

func runDay13() {
	day13.RunDay13()
}

func runDay14() {
	day14.RunDay14()
}
