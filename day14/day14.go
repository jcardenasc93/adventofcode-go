package day14

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay14() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			processLine(line)
		}
	}

	lettersCount := runSteps(10)
	fmt.Println(len(template))
	lmax, max := getMax(&lettersCount)
	lmin, min := getMin(&lettersCount)

	fmt.Println(lmax, max)
	fmt.Println(lmin, min)

	fmt.Println(max - min)
}

var template string
var rules = map[string]string{}

func processLine(line string) {
	if strings.Contains(line, "->") {
		rule := strings.Split(line, "->")
		key := strings.TrimSpace(rule[0])
		value := strings.TrimSpace(rule[1])
		updateRules(key, value)
	} else {
		template = line
	}
}

func updateRules(k string, v string) {
	rules[k] = v
}

func runSteps(n int) map[string]int {
	var outTemplate string
	lettersCount := templateCount()
	outTemplate = template
	for i := 0; i < n; i++ {
		for j := 0; j < len(outTemplate); j++ {
			if j+1 < len(outTemplate) {
				pair := fmt.Sprintf("%s%s", string(outTemplate[j]), string(outTemplate[j+1]))
				ruleVal := applyRule(pair)
				updateLettersCount(&lettersCount, ruleVal)
			}
		}
		outTemplate = template
	}

	return lettersCount
}

func templateCount() map[string]int {
	lettersCount := map[string]int{}
	for _, l := range template {
		lettersCount[string(l)] += 1
	}
	return lettersCount
}

func applyRule(pair string) string {
	rule := rules[pair]
	applied := fmt.Sprintf("%s%s%s", string(pair[0]), rule, string(pair[1]))
	template = strings.Replace(template, pair, applied, 1)
	return rule
}

func updateLettersCount(lettersCount *map[string]int, l string) {
	(*lettersCount)[l] += 1
}

func getMax(lc *map[string]int) (string, int) {
	var max int
	var l string
	for k, v := range *lc {
		if v > max {
			l = k
			max = v
		}
	}
	return l, max
}

func getMin(lc *map[string]int) (string, int) {
	min := math.MaxInt
	var l string
	for k, v := range *lc {
		if v < min {
			l = k
			min = v
		}
	}
	return l, min
}
