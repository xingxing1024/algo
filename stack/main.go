package main

import (
	"container/list"
	"fmt"
)

// 单调栈 柱状图的最大面积
func main() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)
	for stack.Len() > 0 {
		back := stack.Back()
		stack.Remove(back)
		fmt.Println(back.Value.(int))
	}
}
