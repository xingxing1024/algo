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

func inside(x, y int, n, m int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func mazeRoute(x, y int, curPathLen int, visited [100][100]int, n, m int, board [100][100]int) int {
	if x == n-1 && y == m-1 {
		return curPathLen
	}

	visited[x][y] = 1

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	result := 1000000
	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		if !inside(newX, newY, n, m) {
			continue
		}
		// 已经遍历过
		if visited[newX][newY] == 1 {
			continue
		}
		// 障碍物
		if board[newX][newY] == 1 {
			continue
		}
		pathLength := mazeRoute(newX, newY, curPathLen+1, visited, n, m, board)
		if pathLength < result {
			result = pathLength
		}
	}
	visited[x][y] = 0
	return result
}

type Map [4][4]int

func floodFill(x, y int, visited *Map, board *Map, fillValue int) {
	visited[x][y] = fillValue

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	n := len(board)
	m := len(board[0])

	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		if !inside(newX, newY, n, m) {
			continue
		}
		if visited[newX][newY] != 0 {
			continue
		}
		// 障碍物
		if board[newX][newY] == 1 {
			continue
		}
		floodFill(newX, newY, visited, board, fillValue)
	}
}

// todo: bfs学习

func main() {
	//var board [100][100]int
	//nQueen(0, 4, board)

	//var visited [100][100]int
	//var board [100][100]int
	//shortestPath := mazeRoute(0, 0, 0, visited, 4, 4, board)
	//fmt.Println(shortestPath)

	// 种子填充
	//var visited Map
	//var board = Map{
	//	{0, 1, 0, 0},
	//	{0, 1, 0, 0},
	//	{0, 1, 0, 0},
	//	{0, 0, 0, 0},
	//}
	//cnt := 0
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 4; j++ {
	//		if visited[i][j] == 0 && board[i][j] == 0 {
	//			cnt++
	//			floodFill(i, j, &visited, &board, cnt)
	//		}
	//	}
	//}
	//fmt.Println(cnt)
	//fmt.Println(visited)
}
