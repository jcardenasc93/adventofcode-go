package day15

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay15() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	riskMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			processLine(line, &riskMap)
		}
	}
	findLowerRiskPath(&riskMap)
}

func processLine(line string, riskMap *[][]int) {
	levels := []int{}
	for _, l := range line {
		level := utils.ParseStrInt(string(l))
		levels = append(levels, level)
	}
	(*riskMap) = append((*riskMap), levels)
}

func findLowerRiskPath(riskMap *[][]int) {
	pq := make(utils.PriorityQueue, 0)
	visited := []string{}
	directions := [][]int{
		{-1, 0}, // top
		{0, 1},  // right
		{1, 0},  // bottom
		{0, -1}, // left
	}

	risks := map[string]int{}

	item := &utils.Item{
		Value:    "0,0",
		Priority: 0,
		Index:    0,
	}
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		coords := heap.Pop(&pq).(*utils.Item)
		if utils.IsItemIn(visited, coords.Value) {
			continue
		}
		visited = append(visited, coords.Value)
		risks[coords.Value] = coords.Priority
		y, x := getCoordsFromStr(coords.Value)

		if y == len(*riskMap)-1 && x == len(*riskMap)-1 {
			fmt.Println(risks[coords.Value])
			return
		}

		for _, dir := range directions {
			y_, x_ := y+dir[0], x+dir[1]
			if !(utils.IsItemIn(visited, fmt.Sprintf("%d,%d", y_, x_))) {
				if y_ >= 0 && y_ < len(*riskMap) && x_ >= 0 && x_ < len(*riskMap) {
					risk := coords.Priority + (*riskMap)[y_][x_]
					nextStep := &utils.Item{
						Value:    fmt.Sprintf("%d,%d", y_, x_),
						Priority: risk,
						// Index:    pq.Len(),
					}
					heap.Push(&pq, nextStep)
				}
			}
		}
	}
}

func getCoordsFromStr(s string) (int, int) {
	split := strings.Split(s, ",")
	y := utils.ParseStrInt(split[0])
	x := utils.ParseStrInt(split[1])
	return y, x
}
