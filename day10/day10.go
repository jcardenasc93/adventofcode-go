package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay10() {
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
	fmt.Println("Error final score", errorScore)
	sort.Ints(completionScores)
	fmt.Println("Completion score", completionScores[len(completionScores)/2])
}

type chunkNode struct {
	value string
	prev  *chunkNode
}

type chunk struct {
	head      *chunkNode
	openTags  []string
	closeTags []string
	lenght    int
}

func createchunk() *chunk {
	return &chunk{
		openTags:  []string{"{", "[", "(", "<"},
		closeTags: []string{"}", "]", ")", ">"},
	}
}

func (c *chunk) add(value string) {
	node := chunkNode{
		value: value,
	}
	c.lenght += 1
	if c.head == nil {
		c.head = &node
	} else {
		node.prev = c.head
		c.head = &node
	}
}

func (c *chunk) peek() string {
	if c.head != nil {
		return c.head.value
	}
	return ""
}

func (c *chunk) remove() {
	c.lenght -= 1
	if c.lenght < 0 {
		c.lenght = 0
		c.head = nil
	} else {
		c.head = c.head.prev
	}
}

func (c *chunk) isValid(value string) bool {
	pairs := map[string]string{
		"}": "{",
		"]": "[",
		">": "<",
		")": "(",
	}

	expected := pairs[value]
	return expected == c.peek()
}

func (c *chunk) autocomplete(score *int) {
	pairs := map[string]string{
		"{": "}",
		"[": "]",
		"<": ">",
		"(": ")",
	}
	for c.lenght > 0 {
		if c.head != nil {
			charCompletion := pairs[c.peek()]
			updateAutocompleteScore(charCompletion, score)
			c.remove()
		}
	}
	completionScores = append(completionScores, *score)
}

func updateErrorScore(char string, score *int) {
	scoreTable := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	*score += scoreTable[char]
}

func updateAutocompleteScore(char string, score *int) {
	scoreTable := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	*score *= 5
	*score += scoreTable[char]
}

var errorScore int
var completionScores []int

func processLine(line string) {
	var autocompleteScore int
	chunkError := false
	chunk := createchunk()
	for _, c := range line {
		cStr := string(c)
		if utils.IsItemIn(chunk.openTags, cStr) == true {
			chunk.add(string(c))
		} else {
			if chunk.isValid(cStr) == true {
				chunk.remove()
			} else {
				updateErrorScore(cStr, &errorScore)
				chunkError = true
				break
			}
		}
	}

	if chunkError == false {
		chunk.autocomplete(&autocompleteScore)
	}
}
