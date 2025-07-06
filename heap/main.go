package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	heapArray := *h
	x := heapArray[len(heapArray)-1]
	*h = heapArray[0 : len(heapArray)-1]
	return x
}

type Person struct {
	Name string
	Age  int
}

type PersonHeap []Person

func (h PersonHeap) Len() int {
	return len(h)
}

func (h PersonHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h PersonHeap) Less(i, j int) bool {
	return h[i].Age > h[j].Age
}

func (h *PersonHeap) Push(x any) {
	*h = append(*h, x.(Person))
}

func (h *PersonHeap) Pop() any {
	heapArray := *h
	x := heapArray[len(heapArray)-1]
	*h = heapArray[0 : len(heapArray)-1]
	return x
}

func main() {

	personHeap := &PersonHeap{
		Person{
			Name: "peter",
			Age:  10,
		},
		Person{
			Name: "tom",
			Age:  20,
		},
	}
	heap.Init(personHeap)
	heap.Push(personHeap, Person{
		Name: "tina",
		Age:  15,
	})
	for personHeap.Len() > 0 {
		fmt.Println(heap.Pop(personHeap))
	}

	//intHeap := &IntHeap{1, 4, 5, 3}
	//heap.Init(intHeap)
	//heap.Push(intHeap, 100)
	//
	//fmt.Println(*intHeap)
	//for intHeap.Len() > 0 {
	//	fmt.Println(heap.Pop(intHeap))
	//}
}
