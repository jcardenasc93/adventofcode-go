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

func calcIntAvg(data []int) int64 {
	var sum int
	for _, d := range data {
		sum += d
	}
	avg := float64(sum) / float64(len(data))
	return int64(math.Round(avg))
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
	minFuel := math.MaxFloat64
	for _, f := range fuels {
		if f < minFuel {
			minFuel = f
		}
	}
	fmt.Println(minFuel)
}

func getDistance(crabs *[]int, value int) float64 {
	var totalFuel float64
	for _, crab := range *crabs {
		fuel := math.Abs(float64(crab) - float64(value))
		totalFuel += fuel
	}
	return totalFuel
}
