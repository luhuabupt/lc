package main

func main() {

}

type ListNode struct {
	v    int
	next *ListNode
}

// O(n)
// 1. reverse half
// 2. p1 0, p2 half
func reListNode_(head *ListNode) bool {
	arr := []int{}
	for p := head; p != nil; p = p.next {
		arr = append(arr, p.v)
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}
