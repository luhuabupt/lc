//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
// 示例 1：
//
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//
//
// 示例 2：
//
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//
//
// 提示：
//
//
// 2 <= timePoints <= 2 * 104
// timePoints[i] 格式为 "HH:MM"
//
// Related Topics 数组 数学 字符串 排序
// 👍 134 👎 0

package main

import (
	"strconv"
	"strings"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findMinDifference(timePoints []string) int {
	ex := map[int]bool{}
	for _, v := range timePoints {
		xa := strings.Split(v, ":")
		h, _ := strconv.Atoi(xa[0])
		m, _ := strconv.Atoi(xa[1])
		if ex[h*60+m] {
			return 0
		}
		ex[h*60+m] = true
	}
	arr := []int{}
	for i := 0; i < 24*60; i++ {
		if ex[i] {
			arr = append(arr, i)
		}
	}

	ans := 60 * 24
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			if arr[0]+24*60-arr[len(arr)-1] < ans {
				ans = arr[0] + 24*60 - arr[len(arr)-1]
			}
		} else if arr[i]-arr[i-1] < ans {
			ans = arr[i] - arr[i-1]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
