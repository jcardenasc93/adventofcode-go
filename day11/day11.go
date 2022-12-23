package day11

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay11() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	octopuses := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			octopuses = append(octopuses, line)
		}
	}
	runSteps(octopuses, 100)
}

type point struct {
	x int
	y int
}

type octoLevels struct {
	octs       [10][10]int
	flashes    []point
	flashCount int
}

func runSteps(octopuses []string, n int) {
	ol := octoLevels{}
	for i := 0; i < len(octopuses); i++ {
		for j, v := range octopuses[i] {
			ol.octs[i][j] = utils.ParseStrInt(string(v))
		}
	}

	for i := 0; i < n; i++ {
		ol.step()
	}

	for _, i := range ol.octs {
		for _, j := range i {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
	fmt.Println("total flashes:", ol.flashCount)
}

func (ol *octoLevels) step() {
	ol.flashes = []point{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			p := point{i, j}
			ol.increaseLevel(p)
		}
	}
	// processFlashes
	for _, f := range ol.flashes {
		ol.processFlash(&f)
	}
}

func (ol *octoLevels) increaseLevel(p point) {
	ol.octs[p.x][p.y] += 1
	if ol.octs[p.x][p.y] > 9 {
		if utils.IsItemIn(ol.flashes, p) == false {
			ol.flashes = append(ol.flashes, p)
			ol.flashCount += 1
		}
	}
}

func (ol *octoLevels) processFlash(flash *point) {
	directions := map[string][]int{
		"topLeft":     {-1, -1},
		"top":         {0, -1},
		"topRight":    {1, -1},
		"left":        {-1, 0},
		"right":       {1, 0},
		"bottomLeft":  {-1, 1},
		"bottom":      {0, 1},
		"bottomRight": {1, 1},
	}

	for _, d := range directions {
		i := flash.x + d[0]
		j := flash.y + d[1]
		if i >= 0 && i < 10 && j >= 0 && j < 10 {
			p := point{i, j} // add cond breakp for flash.x = 3 && flash.y == 9
			prevFlash := utils.IsItemIn(ol.flashes, p)
			if prevFlash == false && ol.octs[i][j] != 0 {
				ol.increaseLevel(p)
			}
			newFlash := utils.IsItemIn(ol.flashes, p)
			if prevFlash == false && newFlash == true {
				ol.processFlash(&p)
			}
		}
	}
	ol.octs[flash.x][flash.y] = 0
}
