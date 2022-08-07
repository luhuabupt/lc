package main

import (
	"fmt"
	"sort"
	"strings"
)

// reorder-data-in-log-files 重新排列日志文件  2022-05-03 08:34:00
func main() {
	//fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
	//fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}))
	fmt.Println(reorderLogFiles([]string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reorderLogFiles(logs []string) []string {
	isNum := func(v byte) int {
		if v >= '0' && v <= '9' {
			return 1
		}
		return 0
	}

	sort.SliceStable(logs, func(i, j int) bool {
		sa := strings.Split(logs[i], " ")
		sb := strings.Split(logs[j], " ")
		if isNum(sa[1][0]) != isNum(sb[1][0]) {
			return isNum(sa[1][0]) < isNum(sb[1][0])
		}
		if isNum(sa[1][0]) == 1 {
			return i < j
		}
		if logs[i][len(sa[0]):] != logs[j][len(sb[0]):] {
			return logs[i][len(sa[0]):] < logs[j][len(sb[0]):]
		}
		return sa[0] < sb[0]
	})
	return logs
}

//leetcode submit region end(Prohibit modification and deletion)
