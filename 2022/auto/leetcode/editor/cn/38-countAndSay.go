package main

import "fmt"

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countAndSay(n int) string {
	itob := func(v int) (re []byte) {
		for v > 0 {
			re = append(re, byte('0'+v%10))
			v /= 10
		}
		for i := 0; i < len(re)/2; i++ {
			re[i], re[len(re)-i] = re[len(re)-i], re[i]
		}
		return
	}

	pre, cur := []byte{'1'}, []byte{'1'}
	for i := 2; i <= n; i++ {
		pre, cur = cur, []byte{}
		v, cnt := pre[0], 1
		for i := 1; i < len(pre); i++ {
			if pre[i] == v {
				cnt++
			} else {
				cur = append(cur, itob(cnt)...)
				cur = append(cur, v)
				v, cnt = pre[i], 1
			}
		}
		cur = append(cur, itob(cnt)...)
		cur = append(cur, v)
	}

	return string(cur)
}

//leetcode submit region end(Prohibit modification and deletion)
