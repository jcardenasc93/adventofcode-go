package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

var uniqueDigitsCounter int

func RunDay8() {
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
		processInput(input)
	}
	fmt.Println(uniqueDigitsCounter)
}

func processInput(input string) {
	if input != "" {
		signals := strings.Split(strings.Split(input, " | ")[1], " ")
		for _, s := range signals {
			sLen := len(s)
			if sLen == 2 || sLen == 4 || sLen == 3 || sLen == 7 {
				uniqueDigitsCounter += 1
			}
		}
	}
}
