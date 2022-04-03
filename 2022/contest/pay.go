package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(coopDevelop([][]int{{1, 2, 3}, {3}, {2, 4}}))
	fmt.Println(coopDevelop([][]int{{3}, {6}}))
}

func coopDevelop(skills [][]int) (ans int) {
	dp := map[string]int{}
	M := int(1e9) + 7
	sort.Slice(skills, func(i, j int) bool {
		return len(skills[i]) < len(skills[j])
	})

	for i, v := range skills {
		cant, end := 0, ""
		for x := 1; x < 1<<len(v); x++ {
			t := []string{}
			for j := 0; j < len(v); j++ {
				if x>>j&1 == 1 {
					t = append(t, strconv.Itoa(v[j]))
				}
			}
			st := strings.Join(t, "_")
			cant += dp[st]
			if x+1 == 1<<len(v) {
				end = st
			}
		}
		if i > cant {
			ans += (i - cant)
			ans %= M
		}
		dp[end]++
	}

	return ans
}
