package sortalgo

func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				Swap(data, i, j)
			}
		}
	}
}
