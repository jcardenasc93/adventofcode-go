package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay7() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	processInput(input)
}

func processInput(input string) {
	crabs := utils.ParseStringToIntSlice(input)
	min := math.MaxInt
	max := 0
	for _, crab := range crabs {
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}
	}
	var fuels []float64
	for i := min; i <= max; i++ {
		fuels = append(fuels, getDistance(&crabs, i))
	}
	var fuels2 []float64
	for i := min; i <= max; i++ {
		fuels2 = append(fuels2, getDistanceCrabEng(&crabs, i))
	}

	fmt.Println("Part one: ", getMinFuel(&fuels))
	fmt.Println("Part two: ", getMinFuel(&fuels2))
}

func getMinFuel(fuels *[]float64) float64 {
	minFuel := math.MaxFloat64
	for _, f := range *fuels {
		if f < minFuel {
			minFuel = f
		}
	}
	return minFuel
}

func getDistance(crabs *[]int, value int) float64 {
	var totalFuel float64
	for _, crab := range *crabs {
		fuel := math.Abs(float64(crab) - float64(value))
		totalFuel += fuel
	}
	return totalFuel
}

func getDistanceCrabEng(crabs *[]int, value int) float64 {
	var totalFuel float64
	var fuel float64
	for _, crab := range *crabs {
		fuel = float64(0)
		bound := math.Abs(float64(crab) - float64(value))
		for i := 0; float64(i) <= bound; i++ {
			fuel += float64(i)
		}
		totalFuel += fuel
	}
	return totalFuel
}
