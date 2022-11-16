package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDay6() {
	readInputFile()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func readInputFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	processLanternFish(input)
}

func processLanternFish(input string) {
	var lanternFishState []int
	for _, fish := range strings.Split(input, ",") {
		fishInt, err := strconv.Atoi(fish)
		checkError(err)
		lanternFishState = append(lanternFishState, fishInt)
	}
	days := 80
	for i := 1; i <= days; i++ {
		nextDay(&lanternFishState)
	}
	fmt.Println("After", days, "days there is a total of ", len(lanternFishState))
}

func nextDay(fish *[]int) {
	for i := range *fish {
		f := &(*fish)[i]
		*f -= 1
		if *f < 0 {
			*f = 6
			*fish = append(*fish, 8)
		}
	}
}
