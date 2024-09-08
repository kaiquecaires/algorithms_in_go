package classphotos

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	back := "red"

	if blueShirtHeights[0] > redShirtHeights[0] {
		back = "blue"
	}

	for i, rHeight := range redShirtHeights {
		if back == "red" {
			if rHeight <= blueShirtHeights[i] {
				return false
			}
		} else {
			if blueShirtHeights[i] <= rHeight {
				return false
			}
		}
	}

	return true
}
