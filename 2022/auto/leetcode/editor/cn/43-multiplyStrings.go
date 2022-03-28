package main

import "fmt"

func main() {
	fmt.Println(multiply("99", "72"))
	fmt.Println(multiply("123", "456"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(aa string, bb string) string {
	a := []byte(aa)
	b := []byte(bb)

	ans := make([]int, len(a)*len(b))
	for i, v := range a {
		for j, x := range b {
			t := int(v-'0') * int(x-'0')
			ans[len(a)-1-i+len(b)-1-j] += t
		}
	}

	//fmt.Println(ans)

	re := []byte{}
	ca := 0
	for _, v := range ans {
		re = append(re, byte('0'+(v+ca)%10))
		ca = (ca + v) / 10
	}
	re = append(re, byte('0'+ca))

	for i := len(re) - 1; i >= 0 && re[i] == '0'; i-- {
		re = re[:len(re)-1]
	}

	for i := 0; i < len(re)/2; i++ {
		re[i], re[len(re)-1-i] = re[len(re)-1-i], re[i]
	}

	if len(re) == 0 {
		return "0"
	}

	return string(re)
}

//leetcode submit region end(Prohibit modification and deletion)
