package main

import (
	sortalgo "algo/sort_algo"
	"fmt"
)

func main() {
	data := []int{1, 54, 3, 2, 45}
	sortalgo.BubbleSort(data)
	fmt.Println(data)
}
