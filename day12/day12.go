package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	pathsCount = 0
	visited = make(map[string]int)
	findPaths2("start")
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
	if utils.IsLower(cave) {
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

func findPaths2(cave string) {
	if cave == "end" {
		pathsCount += 1
		return
	}
	if utils.IsLower(cave) {
		visited[cave] += 1
		if reachMaxVisits(&visited, cave) {
			return
		}
	}
	for _, c := range pathsGraph.adjList[cave] {
		if c != "start" {
			findPaths2(c)
		}
	}
	visited[cave] -= 1
}

func reachMaxVisits(visited *map[string]int, cave string) bool {
	maxLower := 2
	var twiceCount int
	for _, v := range *visited {
		if v > 1 {
			twiceCount += 1
		}
		if v > maxLower {
			(*visited)[cave] -= 1
			return true
		}
	}

	if twiceCount > 1 {
		(*visited)[cave] -= 1
		return true
	}

	return false
}
