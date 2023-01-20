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

	pairsCount := countPairs()
	charCount := runSteps(40, &pairsCount)
	fmt.Println(pairsCount)
	fmt.Println(charCount)
	lmax, max := getMax(&charCount)
	lmin, min := getMin(&charCount)

	fmt.Println(lmax, max)
	fmt.Println(lmin, min)

	fmt.Println(max - min)
}

var template string
var rules = map[string]string{}

func processLine(line string) {
	if strings.Contains(line, "->") {
		rule := strings.Split(line, " -> ")
		key := rule[0]
		value := rule[1]
		rules[key] = value
	} else {
		template = line
	}
}

func runSteps(n int, pairsCount *map[string]int) map[string]int {
	var charCount map[string]int
	for i := 0; i < n; i++ {
		insertions := []map[string]int{}
		charCount = map[string]int{}
		for p := range *pairsCount {
			applyRule(p, pairsCount, &insertions)
		}

		for _, i := range insertions {
			for k, v := range i {
				newPairs := strings.Split(k, ",")
				leftPair, rightPair := newPairs[0], newPairs[1]
				// Each time a new char is inserted two new pairs are created
				// and at the same time the initial pair dissapears
				(*pairsCount)[leftPair] += v
				(*pairsCount)[rightPair] += v

				// Only left new pair is considered to count char in order to avoid double count
				charCount[string(leftPair[0])] += v
				charCount[string(leftPair[1])] += v
			}
		}
	}

	// Adds the left value for the last char in the template
	charCount[string(template[len(template)-1])] += 1
	return charCount
}

func countPairs() map[string]int {
	pairsCount := map[string]int{}
	for i := 0; i < len(template); i++ {
		if i+1 < len(template) {
			pair := fmt.Sprintf("%s%s", string(template[i]), string(template[i+1]))
			pairsCount[pair] += 1
		}
	}
	return pairsCount
}

func applyRule(pair string, pairCount *map[string]int, insertions *[]map[string]int) {
	rule, ok := rules[pair]
	if ok {
		leftPair := fmt.Sprintf("%s%s", string(pair[0]), rule)
		rightPair := fmt.Sprintf("%s%s", rule, string(pair[1]))

		insertion := map[string]int{fmt.Sprintf("%s,%s", leftPair, rightPair): (*pairCount)[pair]}
		(*insertions) = append(*insertions, insertion)

		// Each time a new char is inserted two new pairs are created
		// and at the same time the initial pair dissapears
		(*pairCount)[pair] -= (*pairCount)[pair]
	}
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
