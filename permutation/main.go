package main

import "fmt"

// Permutation 全排列
func Permutation(curPermu []int, nums []int, visited map[int]bool) {
	if len(curPermu) == len(nums) {
		fmt.Println(curPermu)
		return
	}

	for _, n := range nums {
		if visited[n] {
			continue
		}
		visited[n] = true
		Permutation(append(curPermu, n), nums, visited)
		visited[n] = false
	}
}

func main() {
	visited := make(map[int]bool)
	Permutation(nil, []int{1, 2, 3}, visited)
}
