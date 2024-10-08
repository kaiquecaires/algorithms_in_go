package hassinglecycle

func HasSingleCycle(array []int) bool {
	jumpTo := JumpTo(array, 0)
	jumps := 0

	for array[jumpTo] != 0 {
		prev := jumpTo
		jumpTo = JumpTo(array, jumpTo)
		array[prev] = 0
		jumps++
	}

	return jumps == len(array)
}

func JumpTo(array []int, idx int) int {
	n := len(array)
	jump := (idx + array[idx]) % n
	if jump < 0 {
		jump += n
	}
	return jump
}
