package selectionsort

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		smIdx := i

		for j := i; j < len(array); j++ {
			if array[j] < array[smIdx] {
				smIdx = j
			}
		}

		array[i], array[smIdx] = array[smIdx], array[i]
	}

	return array
}
