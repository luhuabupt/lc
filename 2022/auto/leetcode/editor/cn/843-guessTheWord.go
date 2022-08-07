package main

func main() {

}

type Master struct {
}

func (this *Master) Guess(word string) int { return 0 }

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */

func findSecretWord(w []string, m *Master) {
	chk := func(a, b string) (ans int) {
		for i := 0; i < 6; i++ {
			if a[i] != b[i] {
				ans++
			}
		}
		return
	}

	for i := 0; i < 10; i++ {
		r := m.Guess(w[i])
		if r == 6 {
			break
		}

		nw := []string{}
		for i := 1; i < len(w); i++ {
			if chk(w[0], w[i]) == r {
				nw = append(nw, w[i])
			}
		}
		w = nw
	}
}

//leetcode submit region end(Prohibit modification and deletion)
