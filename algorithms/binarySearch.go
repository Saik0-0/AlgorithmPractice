package algorithms

func BinarySearch(array []int, item int) (elemInd int, found bool) {
	begin := 0
	end := len(array) - 1

	for begin <= end {
		elemInd = (begin + end) / 2
		elem := array[elemInd]
		if elem == item {
			return elemInd, true
		}
		if elem > item {
			end = elemInd - 1
		} else {
			begin = elemInd + 1
		}
	}
	return -1, false
}

func StupidSearch(array []int, item int) (_ int, found bool) {
	for i := range len(array) {
		if array[i] == item {
			return i, true
		}
	}
	return -1, false
}
