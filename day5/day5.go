package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	xRange []int
	yRange []int
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

func initSegment(ends []string) Segment {
	var minX int
	var minY int
	var maxX int
	var maxY int

	// Get max & min of x
	x1, err := strconv.Atoi(strings.Split(ends[0], ",")[0])
	checkError(err)
	x2, err := strconv.Atoi(strings.Split(ends[1], ",")[0])
	checkError(err)
	maxX, minX = getMaxMin(x1, x2)

	// Get max & min of y
	y1, err := strconv.Atoi(strings.Split(ends[0], ",")[1])
	checkError(err)
	y2, err := strconv.Atoi(strings.Split(ends[1], ",")[1])
	checkError(err)
	maxY, minY = getMaxMin(y1, y2)

	xrange := []int{}
	yrange := []int{}

	for i := minX; i <= maxX; i++ {
		xrange = append(xrange, i)
	}

	for i := minY; i <= maxY; i++ {
		yrange = append(yrange, i)
	}

	segment := Segment{
		xRange: xrange,
		yRange: yrange,
	}
	return segment
}

func readInputFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	segments := []Segment{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputSegment := strings.Split(line, " -> ")
		segments = append(segments, initSegment(inputSegment))
	}
	fmt.Println(segments)
}
