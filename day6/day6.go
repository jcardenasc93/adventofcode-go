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
	lanternFishState := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fish := range strings.Split(input, ",") {
		fishInt, err := strconv.Atoi(fish)
		checkError(err)
		lanternFishState[fishInt] += 1
	}
	days := 256
	for i := 0; i < days; i++ {
		d0 := lanternFishState[0]
		lanternFishState[0] = 0
		for j := 1; j < 9; j++ {
			lanternFishState[j-1] = lanternFishState[j]
			lanternFishState[j] = 0
		}
		lanternFishState[8] = d0
		lanternFishState[6] += d0
	}
	count := 0
	for _, fish := range lanternFishState {
		count += fish
	}
	fmt.Println("After", days, "days there is a total of ", count)
}
