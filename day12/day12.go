package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay12() {
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
			pathsGraph.initGraph(line)
		}
	}
	fmt.Println(pathsGraph.adjList)
	findPaths("start")
	fmt.Println(pathsCount)
}

var pathsGraph = Graph{
	adjList: map[string][]string{},
}

type Graph struct {
	adjList map[string][]string
}

func (g *Graph) initGraph(input string) {
	caves := strings.Split(input, "-")
	c0, c1 := caves[0], caves[1]
	g.adjList[c0] = append(g.adjList[c0], c1)
	g.adjList[c1] = append(g.adjList[c1], c0)
}

var pathsCount int
var visited = make(map[string]int)

func findPaths(cave string) {
	if cave == "end" {
		pathsCount += 1
		return
	}
	if isLower(cave) {
		visited[cave] += 1
		if visited[cave] > 1 {
			visited[cave] -= 1
			return
		}
	}
	for _, c := range pathsGraph.adjList[cave] {
		if c != "start" {
			findPaths(c)
		}
	}
	visited[cave] -= 1
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
