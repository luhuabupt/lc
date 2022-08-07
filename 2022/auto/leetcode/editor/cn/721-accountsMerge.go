package main

import "fmt"

// accounts-merge 账户合并  2022-01-13 22:38:34
func main() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func accountsMerge(accounts [][]string) [][]string {
	fa := map[string]int{}
	for i, arr := range accounts {
		for _, v := range arr {
			if _, ok := fa[v]; !ok {
				fa[v] = i
			} else {
				if accounts[fa[v]][0] == arr[0] {

				}
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
