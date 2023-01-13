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
	var coords []map[string]int
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if strings.Contains(line, "fold") {
				line := strings.Split(line, " ")
				splitLine := strings.Split(line[2], "=")
				fold = append(fold, splitLine...)
				break
			}
			coords = append(coords, getCoords(line))
		}
	}
	applyFold(&coords, fold)
}

func getCoords(cord string) map[string]int {
	mapCord := make(map[string]int)
	cords := strings.Split(cord, ",")
	mapCord["x"] = utils.ParseStrInt(cords[0])
	mapCord["y"] = utils.ParseStrInt(cords[1])
	return mapCord
}

func applyFold(coords *[]map[string]int, fold []string) {
	foldDir := fold[0]
	foldLine := utils.ParseStrInt(fold[1])
	pointsCount := 0

	for _, coord := range *coords {
		if coord[foldDir] < foldLine {
			pointsCount += 1
		} else {
			if !(isOverlaping(coords, coord, foldDir, foldLine)) {
				pointsCount += 1
			}
		}
	}
	fmt.Println(pointsCount)
}

func isOverlaping(coords *[]map[string]int, coord map[string]int, foldDir string, foldLine int) bool {
	partialCord := map[string]int{}
	partialCord[foldDir] = 2*foldLine - coord[foldDir]
	if foldDir == "x" {
		partialCord["y"] = coord["y"]
	} else {
		partialCord["x"] = coord["x"]
	}
	return utils.IsMapIn(coords, partialCord)

}
