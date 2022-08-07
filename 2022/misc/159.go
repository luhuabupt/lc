package main

func lengthOfLongestSubstringTwoDistinct(s string) (ans int) {
	cnt := map[uint8]int{}
	no := 0
	for i, j := 0, 0; j < len(s); j++ {
		cnt[s[j]]++
		if cnt[s[j]] == 1 {
			no++
		}

		for ; no > 2; i++ {
			cnt[s[i]]--
			if cnt[s[i]] == 0 {
				no--
			}
		}

		if j-i > ans {
			ans = j - i
		}
	}

	return ans + 1
}
