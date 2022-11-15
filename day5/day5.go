package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lenX int
var lenY int

func RunDay5() {
	readInputFile()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

type Segment struct {
	minX int
	minY int
	maxX int
	maxY int
}

func getMaxMin(val1 int, val2 int) (int, int) {
	var min int
	var max int
	if val1 <= val2 {
		min = val1
		max = val2
	} else {
		min = val2
		max = val1
	}
	return max, min
}

func isValidSegment(segment []string) (bool, []int, []int) {
	x1, err := strconv.Atoi(strings.Split(segment[0], ",")[0])
	checkError(err)
	x2, err := strconv.Atoi(strings.Split(segment[1], ",")[0])
	checkError(err)

	// Get max & min of y
	y1, err := strconv.Atoi(strings.Split(segment[0], ",")[1])
	checkError(err)
	y2, err := strconv.Atoi(strings.Split(segment[1], ",")[1])
	checkError(err)

	return (x1 == x2 || y1 == y2), []int{x1, x2}, []int{y1, y2}
}

func initSegment(xVals []int, yVals []int) Segment {
	var minX int
	var minY int
	var maxX int
	var maxY int

	// Get max & min of x
	maxX, minX = getMaxMin(xVals[0], xVals[1])

	// Get max & min of y
	maxY, minY = getMaxMin(yVals[0], yVals[1])

	segment := Segment{
		minX: minX,
		minY: minY,
		maxX: maxX,
		maxY: maxY,
	}

	// Update matrix length
	if maxY > lenY {
		lenY = maxY + 1
	}
	if maxX > lenX {
		lenX = maxX + 1
	}

	return segment
}

func readInputFile() {
	// file, err := os.Open("file.txt.bp")
	file, err := os.Open("file.txt.bp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	segments := []Segment{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputSegment := strings.Split(line, " -> ")
		segmentValid, xVals, yVals := isValidSegment(inputSegment)
		if segmentValid == true {
			segments = append(segments, initSegment(xVals, yVals))
		}
	}
	processSegments(segments)
}

func processSegments(segments []Segment) {
	fmt.Println(len(segments))
	fmt.Println(countLines(&segments))
}

func countLines(lines *[]Segment) int {
	var count int
	size, _ := getMaxMin(lenX, lenY)
	linesMatrix := make([][]int, size)
	for i := range linesMatrix {
		linesMatrix[i] = make([]int, size)
	}

	for _, line := range *lines {
		for j := line.minY; j <= line.maxY; j++ {
			for i := line.minX; i <= line.maxX; i++ {
				linesMatrix[j][i] += 1
			}
		}
	}

	for j, row := range linesMatrix {
		for i := range row {
			if linesMatrix[j][i] >= 2 {
				count += 1
			}
		}
	}
	return count
}
