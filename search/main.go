package main

import "fmt"

func isValidBoard(boardSize int, board [100][100]int) bool {
	rowCnt := make(map[int]int, 0)
	colCnt := make(map[int]int, 0)
	sumCnt := make(map[int]int, 0)
	subCnt := make(map[int]int, 0)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 1 {
				rowCnt[i]++
				colCnt[j]++
				sumCnt[i+j]++
				subCnt[i-j]++
				if rowCnt[i] > 1 {
					return false
				}
				if colCnt[j] > 1 {
					return false
				}
				if sumCnt[i+j] > 1 {
					return false
				}
				if subCnt[i-j] > 1 {
					return false
				}
			}
		}
	}
	return true
}

// 搜索算法 dfs bfs
func nQueen(row int, boardSize int, board [100][100]int) {
	if row == boardSize {
		fmt.Println("======")
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				fmt.Print(board[i][j])
			}
			fmt.Println()
		}
		return
	}
	for j := 0; j < boardSize; j++ {
		board[row][j] = 1
		if !isValidBoard(boardSize, board) {
			board[row][j] = 0
			continue
		}
		nQueen(row+1, boardSize, board)
		board[row][j] = 0
	}
}

func main() {
	var board [100][100]int
	nQueen(0, 4, board)
}
