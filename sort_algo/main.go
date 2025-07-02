package main

import "fmt"

func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

// bubbleSort 冒泡排序
func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				Swap(data, i, j)
			}
		}
	}
}

// quickSort 快速排序
func quickSort(data []int, l, r int) {
	if l == r {
		return
	}
	if l > r {
		return
	}

	mid := data[r]
	i := l // 小于i的下标的均小于mid 大于等于i的均大于等于mid
	for j := l; j <= r; j++ {
		if data[j] < mid {
			Swap(data, i, j)
			i++
		}
	}
	Swap(data, i, r)
	quickSort(data, l, i-1)
	quickSort(data, i+1, r)
}

// mergeSort 归并排序
func mergeSort(data []int, l, r int) {
	if l == r {
		return
	}
	if l > r {
		return
	}

	mid := (l + r) / 2
	mergeSort(data, l, mid)
	mergeSort(data, mid+1, r)

	sortedArr := make([]int, 0)
	leftIdx := l
	rightIdx := mid + 1
	for leftIdx <= mid && rightIdx <= r {
		if data[leftIdx] < data[rightIdx] {
			sortedArr = append(sortedArr, data[leftIdx])
			leftIdx++
		} else {
			sortedArr = append(sortedArr, data[rightIdx])
			rightIdx++
		}
	}

	// 清楚残留
	for leftIdx <= mid {
		sortedArr = append(sortedArr, data[leftIdx])
		leftIdx++
	}
	for rightIdx <= r {
		sortedArr = append(sortedArr, data[rightIdx])
		rightIdx++
	}

	// 重新赋值
	for i := l; i <= r; i++ {
		data[i] = sortedArr[i-l]
	}
	return
}

func main() {
	data := []int{1, 54, 3, 2, 45}
	//bubbleSort(data)
	//quickSort(data, 0, len(data)-1)
	mergeSort(data, 0, len(data)-1)
	fmt.Println(data)
}
