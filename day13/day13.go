package day13

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay13() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fold := []string{}
	var origamiPoints []map[string]int
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if strings.Contains(line, "fold") {
				line := strings.Split(line, " ")
				splitLine := strings.Split(line[2], "=")
				fold = append([]string{}, splitLine...)
				origamiPoints = applyFold(&origamiPoints, fold)
				// break
			} else {
				origamiPoints = append(origamiPoints, getCoords(line))
			}
		}
	}
	printOrigami(&origamiPoints)
}

var maxX, maxY int

func getCoords(cord string) map[string]int {
	mapCord := make(map[string]int)
	cords := strings.Split(cord, ",")
	x_ := utils.ParseStrInt(cords[0])
	y_ := utils.ParseStrInt(cords[1])
	mapCord["x"] = x_
	mapCord["y"] = y_
	return mapCord
}

func applyFold(coords *[]map[string]int, fold []string) []map[string]int {
	foldDir := fold[0]
	foldLine := utils.ParseStrInt(fold[1])
	updateMax(foldLine, foldDir)
	var leftPoints []map[string]int

	for _, coord := range *coords {
		if coord[foldDir] < foldLine {
			leftPoints = append(leftPoints, coord)
		} else {
			point, overlaping := isOverlaping(&leftPoints, coord, foldDir, foldLine)
			if !overlaping {
				leftPoints = append(leftPoints, point)
			}
		}
	}
	return leftPoints
}

func updateMax(value int, axis string) {
	if axis == "x" {
		maxX = value
	} else {
		maxY = value
	}
}

func isOverlaping(coords *[]map[string]int, coord map[string]int, foldDir string, foldLine int) (map[string]int, bool) {
	partialCord := map[string]int{}
	partialCord[foldDir] = 2*foldLine - coord[foldDir]
	if foldDir == "x" {
		partialCord["y"] = coord["y"]
	} else {
		partialCord["x"] = coord["x"]
	}

	return partialCord, utils.IsMapIn(coords, partialCord)
}

func printOrigami(origamiPoints *[]map[string]int) {
	origami := make([][]string, maxY)
	var pointsCount int
	for i := range origami {
		origami[i] = make([]string, maxX)
	}

	for _, point := range *origamiPoints {
		origami[point["y"]][point["x"]] = "#"
		pointsCount += 1
	}

	fmt.Println(pointsCount)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if origami[y][x] == "" {
				fmt.Printf(" ")
			}
			fmt.Printf("%s", origami[y][x])
		}
		fmt.Println()
	}
}
