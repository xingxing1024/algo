package main

import (
	"container/list"
	"fmt"
	"github.com/samber/lo"
)

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

type WaterNode struct {
	va int
	vb int
}

func getNextState(cupa, cupb int, a, b int) []WaterNode {
	return []WaterNode{
		// 分别倒满
		{
			va: a,
			vb: cupb,
		},
		{
			va: cupa,
			vb: b,
		},
		// 分别清空
		{
			va: 0,
			vb: cupb,
		},
		{
			va: cupa,
			vb: 0,
		},
		// 互相倒
		{
			va: lo.Max([]int{0, cupa + cupb - b}),
			vb: lo.Min([]int{cupa + cupb, b}),
		},
		{
			va: lo.Min([]int{cupa + cupb, a}),
			vb: lo.Max([]int{cupa + cupb - a, 0}),
		},
	}
}

func solveWater(a, b int) bool {
	queue := list.New()
	var visited [100][100]int
	queue.PushBack(WaterNode{0, 0})
	visited[0][0] = 1
	for queue.Len() > 0 {
		// 获取现在的最新状态
		curNode := queue.Front()
		queue.Remove(curNode)
		curWaterNode := curNode.Value.(WaterNode)

		if curWaterNode.va == 4 || curWaterNode.vb == 4 {
			return true
		}

		cupa := curWaterNode.va
		cupb := curWaterNode.vb

		// a倒满
		for _, nextState := range getNextState(cupa, cupb, a, b) {
			if visited[nextState.va][nextState.vb] > 0 {
				continue
			}
			visited[nextState.va][nextState.vb] = 1
			queue.PushBack(nextState)
		}
	}
	return false
}

func toposort(nodeList []int, graph map[int][]int) []int {
	// 计算入度
	degree := make(map[int]int, 0)
	for _, vList := range graph {
		for _, v := range vList {
			degree[v]++
		}
	}

	// 进行topo sort
	topoResult := make([]int, 0)
	queue := list.New()
	for _, u := range nodeList {
		if degree[u] == 0 {
			queue.PushBack(u)
		}
	}
	for queue.Len() > 0 {
		curNode := queue.Front()
		queue.Remove(curNode)
		u := curNode.Value.(int)
		topoResult = append(topoResult, u)

		for _, v := range graph[u] {
			degree[v]--
			if degree[v] == 0 {
				queue.PushBack(v)
			}
		}
	}
	return topoResult
}

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

	//canSolve := solveWater(5, 3)
	//fmt.Println(canSolve)

	//graph := make(map[int][]int, 0)
	//graph[1] = []int{2, 3}
	//graph[2] = []int{3}
	//topoResult := toposort([]int{1, 2, 3}, graph)
	//fmt.Println(topoResult)
}
