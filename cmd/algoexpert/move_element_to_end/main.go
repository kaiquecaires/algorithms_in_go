package moveelementtoend

func MoveElementToEnd(array []int, toMove int) []int {
	left, right := 0, len(array)-1

	for left < right {
		if array[left] == toMove && array[right] != toMove {
			array[left], array[right] = array[right], array[left]
		}

		if array[right] == toMove {
			right--
		}

		if array[left] != toMove {
			left++
		}
	}

	return array
}
