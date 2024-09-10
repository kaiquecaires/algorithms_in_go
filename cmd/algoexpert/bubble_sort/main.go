package bubblesort

func BubbleSort(array []int) []int {
	hadSwap := true

	for hadSwap {
		hadSwap = false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				hadSwap = true
			}
		}
	}

	return array
}
