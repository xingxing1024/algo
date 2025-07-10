package lc_hot100

import "container/list"

type ascStackNode struct {
	Val int
	Idx int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dailyTemperatures(temperatures []int) []int {
	ascStack := list.New() // 声明一个单调递增栈

	result := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for ascStack.Len() > 0 && !(ascStack.Back().Value.(ascStackNode).Val > temperatures[i]) {
			topVal := ascStack.Back()
			ascStack.Remove(topVal)
		}

		if ascStack.Len() == 0 {
			result[i] = 0
		} else {
			topVal := ascStack.Back().Value.(ascStackNode)
			result[i] = abs(i - topVal.Idx)
		}

		ascStack.PushBack(ascStackNode{
			Val: temperatures[i],
			Idx: i,
		})
	}

	return result
}
