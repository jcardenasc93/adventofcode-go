package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDay4() {
	fmt.Println("Input bingo data:")
	readInputFile()
}

type BingoData struct {
	inNums    []string
	boards    map[int][][]int
	boardSize int
	tempRow   []int
	tempCol   [][]int
}

var bingData BingoData

const boardSize int = 5

func inintBingoData(bd *BingoData) {
	bd.tempRow = []int{}
	bd.tempCol = [][]int{}
	bd.boards = map[int][][]int{}

}

func readInputFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inintBingoData(&bingData)
	for scanner.Scan() {
		bingoEvaluator(scanner.Text(), &bingData)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	bingoEvaluate(&bingData)
}

func bingoEvaluator(input string, bingData *BingoData) {
	boardsCounter := len(bingData.boards)
	if input == "" {
		if len(bingData.tempCol) == boardSize {
			// Board numbers are complete. So add board & reset 2d slice
			bingData.boards[boardsCounter] = bingData.tempCol
			bingData.tempCol = [][]int{}
		}

	} else {
		if strings.Index(input, ",") != -1 {
			bingData.inNums = strings.Split(input, ",")
		} else {
			row := strings.Split(input, " ")
			for _, str := range row {
				num, err := strconv.Atoi(str)
				if err == nil {
					bingData.tempRow = append(bingData.tempRow, int(num))
				}
			}
			if len(bingData.tempRow) == boardSize {
				// Adds row to 2d slice & resets row
				bingData.tempCol = append(bingData.tempCol, bingData.tempRow)
				bingData.tempRow = []int{}
			}
		}
	}
}

func bingoEvaluate(bingoData *BingoData) {
	fmt.Println(bingoData)
}
