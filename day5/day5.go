package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var lenX int
var lenY int

func RunDay5() {
	readInputFile(true)
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

type Diagonal struct {
	rangeX []int
	rangeY []int
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
	if maxY >= lenY {
		lenY = maxY + 1
	}
	if maxX >= lenX {
		lenX = maxX + 1
	}

	return segment
}

func initDiagonal(diagonal []string) *Diagonal {
	x1, err := strconv.Atoi(strings.Split(diagonal[0], ",")[0])
	checkError(err)
	x2, err := strconv.Atoi(strings.Split(diagonal[1], ",")[0])
	checkError(err)
	y1, err := strconv.Atoi(strings.Split(diagonal[0], ",")[1])
	checkError(err)
	y2, err := strconv.Atoi(strings.Split(diagonal[1], ",")[1])
	checkError(err)

	// Creates range for x
	var rangex []int
	if x1 <= x2 {
		for i := x1; i <= x2; i++ {
			rangex = append(rangex, i)
		}
	} else {
		for i := x1; i >= x2; i-- {
			rangex = append(rangex, i)
		}
	}

	// Creates range for y
	var rangey []int
	if y1 <= y2 {
		for i := y1; i <= y2; i++ {
			rangey = append(rangey, i)
		}
	} else {
		for i := y1; i >= y2; i-- {
			rangey = append(rangey, i)
		}
	}

	// Update matrix length
	maxX, _ := getMaxMin(x1, x2)
	maxY, _ := getMaxMin(y1, y2)
	if lenX <= maxX {
		lenX = maxX + 1
	}
	if lenY <= maxY {
		lenY = maxY + 1
	}

	if math.Abs(float64(x1)-float64(x2)) == math.Abs(float64(y1)-float64(y2)) {
		d := Diagonal{
			rangeX: rangex,
			rangeY: rangey,
		}
		return &d
	}
	return nil
}

func readInputFile(withDiagonals bool) {
	file, err := os.Open("file.txt.bp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	segments := []Segment{}
	diagonals := []Diagonal{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputSegment := strings.Split(line, " -> ")
		segmentValid, xVals, yVals := isValidSegment(inputSegment)
		if segmentValid == true {
			segments = append(segments, initSegment(xVals, yVals))
		}
		if withDiagonals == true {
			diagonal := initDiagonal(inputSegment)
			if diagonal != nil {
				diagonals = append(diagonals, *diagonal)
			}
		}
	}
	processSegments(segments, diagonals, true)
}

func processSegments(segments []Segment, diagonals []Diagonal, withDiagonals bool) {
	count := 0
	size, _ := getMaxMin(lenX, lenY)
	linesMatrix := make([][]int, size)
	for i := range linesMatrix {
		linesMatrix[i] = make([]int, size)
	}
	processLines(&segments, &linesMatrix, &count)
	if withDiagonals == true {
		processDiagonals(&diagonals, &linesMatrix, &count)
	}
	fmt.Println(count)
}

func processLines(segments *[]Segment, matrix *[][]int, count *int) {
	for _, line := range *segments {
		for j := line.minY; j <= line.maxY; j++ {
			for i := line.minX; i <= line.maxX; i++ {
				(*matrix)[j][i] += 1
				if (*matrix)[j][i] == 2 {
					*count += 1
				}
			}
		}
	}
}

func processDiagonals(diagonals *[]Diagonal, matrix *[][]int, count *int) {
	for _, diagonal := range *diagonals {
		for k := range diagonal.rangeY {
			i := diagonal.rangeX[k]
			j := diagonal.rangeY[k]
			(*matrix)[j][i] += 1
			if (*matrix)[j][i] == 2 {
				*count += 1
			}
		}
	}
}

// func countLines(matrix *[][]int) int {
// 	var count int
// 	for j, row := range *matrix {
// 		for i := range row {
// 			if (*matrix)[j][i] >= 2 {
// 				count += 1
// 			}
// 		}
// 	}
