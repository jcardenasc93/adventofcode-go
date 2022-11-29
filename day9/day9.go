package day9

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay9() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	heightMap := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			heightMap = append(heightMap, line)
		}
	}
	processMap(heightMap)
}

type Point struct {
	x int
	y int
}

type MapProcessor struct {
	heightMap    []string
	lowPoints    []int
	currentPoint *Point
	basins       []Point
	basinsGroups [][]Point
}

func processMap(hMap []string) {
	processor := MapProcessor{
		heightMap: hMap,
		currentPoint: &Point{
			x: 0,
			y: 0,
		},
	}
	riskLevel := 0

	for i := 0; i < len(processor.heightMap); i++ {
		for j := 0; j < len(processor.heightMap[0]); j++ {
			processor.currentPoint.x = j
			processor.currentPoint.y = i
			isLowest := processor.checkLowest()
			if isLowest == true {
				num := utils.ParseStrInt(string(processor.heightMap[i][j]))
				processor.lowPoints = append(processor.lowPoints, num)
				riskLevel += (num + 1)
			}
		}
	}
	fmt.Println("risk level: ", riskLevel)

	// Basins
	basins := []int{}
	maxY := len(processor.heightMap)
	maxX := len(processor.heightMap[0])

	data := processor.dataToDigit()

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			basinSize := lookForBasins(&data, x, y)
			if basinSize > 0 {
				basins = append(basins, basinSize)
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	totalProduct := 1
	for _, b := range basins[:3] {
		totalProduct *= b
	}
	fmt.Println(totalProduct)
}

func (processor *MapProcessor) dataToDigit() [][]int {
	dataInt := [][]int{}
	for i := 0; i < len(processor.heightMap); i++ {
		row := []int{}
		for j := 0; j < len(processor.heightMap[0]); j++ {
			row = append(row, utils.ParseStrInt(string(processor.heightMap[i][j])))
		}
		dataInt = append(dataInt, row)
	}
	return dataInt
}

func (processor *MapProcessor) checkLowest() bool {
	var moves = map[string][]int{
		"top":    {-1, 0},
		"right":  {0, 1},
		"bottom": {1, 0},
		"left":   {0, -1},
	}

	isLowest := false
	movesKeys := processor.getMoves()
	for _, m := range movesKeys {
		x, y := moves[m][1], moves[m][0]
		next := Point{
			x: processor.currentPoint.x + x,
			y: processor.currentPoint.y + y,
		}
		num := utils.ParseStrInt(string(processor.heightMap[next.y][next.x]))
		current := utils.ParseStrInt(string(processor.heightMap[processor.currentPoint.y][processor.currentPoint.x]))
		if current >= num {
			isLowest = false
			break
		}
		isLowest = true
	}
	return isLowest
}

func (processor *MapProcessor) getMoves() []string {
	generalMoves := []string{"top", "right", "bottom", "left"}
	maxX := len(processor.heightMap[0]) - 1
	maxY := len(processor.heightMap) - 1
	if processor.currentPoint.y == 0 {
		generalMoves = utils.DelItem(generalMoves, "top")
	} else if processor.currentPoint.y == maxY {
		generalMoves = utils.DelItem(generalMoves, "bottom")
	}

	if processor.currentPoint.x == 0 {
		generalMoves = utils.DelItem(generalMoves, "left")
	} else if processor.currentPoint.x == maxX {
		generalMoves = utils.DelItem(generalMoves, "right")
	}

	return generalMoves
}

func lookForBasins(data *[][]int, x int, y int) int {
	if isReachable(data, x, y) == false {
		return 0
	}
	// Mark as visited
	(*data)[y][x] = 9
	var moves = [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	count := 1
	for _, m := range moves {
		nextX := x + m[1]
		nextY := y + m[0]
		count += lookForBasins(data, nextX, nextY)
	}
	return count
}

func inBounds(data *[][]int, x int, y int) bool {
	yBound := len(*data)
	xBound := len((*data)[0])
	return x >= 0 && x < xBound && y >= 0 && y < yBound
}

func isReachable(data *[][]int, x int, y int) bool {
	return inBounds(data, x, y) && (*data)[y][x] < 9
}
