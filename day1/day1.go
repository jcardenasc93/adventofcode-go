package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkIncreaseMeasurement1() uint16 {
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
		if checkForError(err) {
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

func checkIncreaseMeasurement2() uint16 {
	var input string
	increasesCount := 0
	lastSum := 0
	scanner := bufio.NewScanner(os.Stdin)
	const slidingWindowLen = 3
	measureSlidingWindow := []int{}
	for {
		scanner.Scan()
		input = scanner.Text()
		measure, err := strconv.Atoi(input)
		if len(measureSlidingWindow) < slidingWindowLen {
			measureSlidingWindow = append(measureSlidingWindow, measure)
		}
		if len(measureSlidingWindow) == slidingWindowLen {
			// Reach max window capacity
			windowSum := sumArray(measureSlidingWindow...)
			if windowSum > lastSum && lastSum != 0 {
				increasesCount += 1
			}
			lastSum = windowSum
			// Free one item from window
			popFirstElement(&measureSlidingWindow)
		}
		if err != nil {
			break
		}
	}
	return uint16(increasesCount)
}

func checkForError(err error) bool {
	if err != nil {
		log.Println("Error reading measure, is not a number")
		return true
	}
	return false
}

func sumArray(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func popFirstElement(slice *[]int) {
	_, *slice = (*slice)[0], (*slice)[1:]
}

func RunDay1() {
	fmt.Println("Running part one")
	fmt.Println("\nPlease enter depht measurement")
	fmt.Println("Total of increases", checkIncreaseMeasurement1())
	fmt.Println("Running part two")
	fmt.Println("Total of increases", checkIncreaseMeasurement2())
}
