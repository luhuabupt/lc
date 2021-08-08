package lib

func BinarySearch(arr []int, v int) int {
	l, r, pos := 0, len(arr)-1, 0
	for l <= r {
		mid := (l + r) / 2
		if v < arr[mid] {
			pos, r = mid, mid-1
		} else {
			pos, l = mid+1, mid+1
		}
	}
	return pos
}
