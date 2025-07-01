package main

import (
	binarysearch "algo/binary_search"
	"fmt"
)

func main() {
	firstResult := binarysearch.BinarySearchFirst([]int{1, 2, 2, 2, 3, 4, 5}, 2)
	fmt.Println(firstResult)

	lastResult := binarysearch.BinarySearchLast([]int{1, 2, 2, 2, 3, 4, 5}, 2)
	fmt.Println(lastResult)
}
