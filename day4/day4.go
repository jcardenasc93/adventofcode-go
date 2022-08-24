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
	winner := false
	var board int
	for i, n := range bingData.inNums {
		m, _ := strconv.Atoi(n)
		if i < boardSize-1 {
			markNums(&bingData, m)
		} else {
			markNums(&bingData, m)
			winner, board = checkWin(bingoData)
		}
		if winner {
			fmt.Printf("Board No %v wins with number %v\n", board, m)
			calcScore(bingoData, m, board)
			break
		}
	}
}

func calcScore(bingData *BingoData, n int, b int) {
	sum := 0
	board := bingData.boards[b]
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	fmt.Println("Score is:", sum*n)

}

func markNums(bingData *BingoData, n int) {
	for _, board := range bingData.boards {
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if board[i][j] == n {
					board[i][j] = -1
				}
			}
		}

	}
}

func checkWin(bingData *BingoData) (bool, int) {
	gotWinner := false
	boardn := -1

	// Check cols
	for b := 0; b < len(bingData.boards); b++ {
		board := bingData.boards[b]
		for i := 0; i < boardSize; i++ {
			if board[i][0] == -1 {
				gotWinner = true
				boardn = b
			} else {
				gotWinner = false
				break
			}
		}
	}
	if gotWinner == false {
		// Check rows
		for b := 0; b < len(bingData.boards); b++ {
			board := bingData.boards[b]
			for i := 0; i < boardSize; i++ {
				for j := 0; j < boardSize; j++ {
					if board[i][j] == -1 {
						gotWinner = true
					} else {
						// if one of the nums is not marked continue with the next row
						gotWinner = false
						break
					}
					if j == boardSize-1 && gotWinner == true {
						break
					}
				}
				if gotWinner {
					break
				}
			}
			if gotWinner {
				boardn = b
				break
			}
		}
	}
	return gotWinner, boardn
}
