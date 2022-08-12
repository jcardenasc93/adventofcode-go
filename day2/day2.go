package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x     int
	depth int
}

type positionV2 struct {
	x     int
	depth int
	aim   int
}

func processPart1() {
	var input string
	var pos position
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input for part one:")
	for {
		scanner.Scan()
		input = scanner.Text()
		if input == "" {
			break
		}
		movementHandler(input, &pos)
	}
	fmt.Println("Final position is:", pos.depth*pos.x)
}

func processPart2() {
	var input string
	var pos positionV2
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input for part two:")
	for {
		scanner.Scan()
		input = scanner.Text()
		if input == "" {
			break
		}
		movementHandlerV2(input, &pos)
	}
	fmt.Println("Final position is:", pos.depth*pos.x)
}

func movementHandler(movement string, pos *position) {
	coordinates := strings.Fields(movement)
	movType := coordinates[0]
	value, _ := strconv.Atoi(coordinates[1])
	if movType == "forward" {
		pos.x += value
	} else {
		if movType == "down" {
			pos.depth += value
		} else {
			pos.depth += (-1 * value)
		}
	}
}

func movementHandlerV2(movement string, pos *positionV2) {
	coordinates := strings.Fields(movement)
	movType := coordinates[0]
	value, _ := strconv.Atoi(coordinates[1])
	if movType == "forward" {
		pos.x += value
		pos.depth += (pos.aim * value)
	} else {
		if movType == "down" {
			pos.aim += value
		} else {
			pos.aim += (-1 * value)
		}
	}
}

func RunDay2() {
	processPart1()
	processPart2()
}
