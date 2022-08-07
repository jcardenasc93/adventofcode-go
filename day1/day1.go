package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkIncreaseMeasurement() uint16 {
	var input string
	increasesCount := 0
	lastMeasure := 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input = scanner.Text()
		if input == "" {
			break
		}
		measure, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Error reading measure, is not a number")
			return 0
		}
		if measure > lastMeasure {
			if lastMeasure != 0 {
				increasesCount += 1
			}
		}
		lastMeasure = measure
	}
	return uint16(increasesCount)
}

func RunDay1() {
	fmt.Println("Total of increases", checkIncreaseMeasurement())
}
