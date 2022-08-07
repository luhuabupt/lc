package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func customSortString(order string, s string) string {
	cnt := [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}
	ans := []byte{}
	for _, v := range order {
		for ; cnt[v-'a'] > 0; cnt[v-'a']-- {
			ans = append(ans, byte(v))
		}
	}
	for i, _ := range cnt {
		for ; cnt[i] > 0; cnt[i]-- {
			ans = append(ans, byte(i+'a'))
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
