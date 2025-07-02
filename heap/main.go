package main

import (
	"container/heap"
	"fmt"
)

// IntHeap 实现 heap.Interface 接口
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	// 如果设置 h[i] < h[j] 就是小顶堆，h[i] > h[j] 就是大顶堆
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	// 创建切片
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}

	// 初始化小顶堆
	heap.Init(h)
	fmt.Println(*h) // [0 1 3 6 2 5 7 9 8 4]

	// Pop 元素
	fmt.Println(heap.Pop(h).(int)) // 0

	// Push 元素
	heap.Push(h, 6)
	fmt.Println(*h) // [1 2 3 6 4 5 7 9 8 6]

	for h.Len() != 0 {
		fmt.Printf("%d ", heap.Pop(h).(int)) // 1 2 3 4 5 6 6 7 8 9
	}
}
