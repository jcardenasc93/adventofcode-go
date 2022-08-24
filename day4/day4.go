package day4

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strconv"
	"strings"
)

func RunDay4() {
	fmt.Println("Reading input file...")
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
	bingoEvaluateLast(&bingData)
}

func bingoEvaluator(input string, bingData *BingoData) {
	boardsCounter := len(bingData.boards)
	if input != "" {
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
			if len(bingData.tempCol) == boardSize {
				// Board numbers are complete. So add board & reset 2d slice
				bingData.boards[boardsCounter] = bingData.tempCol
				bingData.tempCol = [][]int{}
			}
		}

	}
}

func bingoEvaluate(bingoData *BingoData) {
	winner := false
	for i, n := range bingData.inNums {
		m, _ := strconv.Atoi(n)
		for j := 0; j < len(bingoData.boards); j++ {
			board := bingoData.boards[j]
			markNums(board, m)
			if i > boardSize-1 {
				winner = checkWin(board, true)
				if winner {
					fmt.Printf("Board No %v wins with number %v\n", j, m)
					calcScore(bingoData, m, j)
					break
				}
			}
		}
		if winner {
			break
		}
	}
}

func bingoEvaluateLast(bingoData *BingoData) {
	winner := false
	var winboards []int
	var lastn int
	for i := 0; i < len(bingoData.inNums); i++ {
		n := bingoData.inNums[i]
		m, _ := strconv.Atoi(n)
		for j := 0; j < len(bingoData.boards); j++ {
			if slices.Contains(winboards, j) == false {
				board := bingoData.boards[j]
				markNums(board, m)
				if i > boardSize-1 {
					winner = checkWin(board, false)
					if winner {
						winboards = append(winboards, j)
						lastn = m
					}
				}
			}
		}
	}
	calcScore(bingoData, lastn, winboards[len(winboards)-1])
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

func markNums(board [][]int, n int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == n {
				board[i][j] = -1
			}
		}

	}
}

func checkWin(board [][]int, stop bool) bool {
	gotWinner := false

	// Check cols
	// for i := 0; i < boardSize; i++ {
	// 	if board[i][0] == -1 {
	// 		gotWinner = true
	// 	} else {
	// 		gotWinner = false
	// 	}
	// }
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[j][i] == -1 {
				gotWinner = true
			} else {
				// if one of the nums is not marked continue with the next col
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
	if gotWinner == false {
		// Check rows
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
	}
	return gotWinner
}
