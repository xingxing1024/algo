package main

import (
	"fmt"
	"testing"
)

var resultList [][]int

func solve(i int, cur []int, nums []int) {
	if i == len(nums) {
		resultList = append(resultList, cur)
		return
	}

	solve(i+1, cur, nums)

	cur = append(cur, nums[i])
	solve(i+1, cur, nums)
	cur = cur[:len(cur)-1]
}

func subsets(nums []int) [][]int {
	resultList = make([][]int, 0)
	solve(0, []int{}, nums)
	return resultList
}

func TestA(t *testing.T) {
	result := subsets([]int{1, 2, 3})
	fmt.Println(result)
}
