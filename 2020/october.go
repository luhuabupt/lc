package main

import (
	"fmt"
	"sort"
)

func minimumDifference_(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	half := sum / 2

	n := len(nums) / 2
	a := make([][]int, n+1)
	b := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		a[i] = []int{}
		b[i] = []int{}
	}

	for i := 0; i < 1<<n; i++ {
		l, v, v2 := 0, 0, 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j]
				v2 += nums[n+j]
				l++
			}
		}
		a[l] = append(a[l], v)
		b[l] = append(b[l], v2)
	}

	for l := 0; l <= n; l++ {
		sort.Ints(a[l])
		sort.Ints(b[l])
	}

	ans := sum - 2*a[n][0]
	if ans < 0 {
		ans = -ans
	}
	//fmt.Println(a, b)

	for l := 0; l <= n; l++ {
		//fmt.Println(a[l], b[n-l])
		for i, j := 0, len(b[n-l])-1; i < len(a[l]); i++ {
			for a[l][i]+b[n-l][j] > half && j > 0 {
				//fmt.Println(i, j)
				ans = minAns(a[l][i]+b[n-l][j], sum, ans)
				j--
			}
			//fmt.Println(i, j)
			ans = minAns(a[l][i]+b[n-l][j], sum, ans)
		}
	}

	return ans
}

func minAns(a, sum, ans int) int {
	d := sum - 2*a
	if d < 0 {
		d = -d
	}
	if d < ans {
		return d
	}
	return ans
}

func minAbsDifference(nums []int, goal int) int {
	// [-2772,6927,4773,-2687,7167,-8995,2940,8869,526]
	//969621127 | 969590451 969589925
	n := len(nums) / 2
	a, b := []int{}, []int{}

	for i := 0; i < 1<<n; i++ {
		v := 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j]
			}
		}
		a = append(a, v)
	}
	for i := 0; i < 1<<(len(nums)-n); i++ {
		v := 0
		for j := 0; j < len(nums)-n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j+n]
			}
		}
		b = append(b, v)
	}

	sort.Ints(a)
	sort.Ints(b)

	//fmt.Println(a)
	//fmt.Println(b)

	ans := ab(a[0] - goal)

	for i, j := 0, len(b)-1; i < len(a); i++ {
		for j > 0 && a[i]+b[j] > goal {
			if ab(a[i]+b[j]-goal) < ans {
				ans = ab(a[i] + b[j] - goal)
			}
			j--
		}
		if ab(a[i]+b[j]-goal) < ans {
			ans = ab(a[i] + b[j] - goal)
		}
	}

	return ans
}

func ab(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 换位置保证nums1先遍历
	if len(nums2) < len(nums1) || len(nums2) == len(nums1) && nums2[len(nums2)-1] < nums1[len(nums1)-1] {
		return findMedianSortedArrays(nums2, nums1)
	}

	i, j, l, cur, prev := 0, 0, 1, 0, 0
	for i < len(nums1) && j < len(nums2) && l <= (len(nums1)+len(nums2)+2)/2 {
		if i < len(nums1) && nums1[i] <= nums2[j] {
			prev, cur = cur, nums1[i]
			i++
		} else {
			prev, cur = cur, nums2[j]
			j++
		}
		l++
	}

	for j < len(nums2) && l <= (len(nums1)+len(nums2)+2)/2 {
		prev, cur = cur, nums2[j]
		j++
		l++
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64(prev) + float64(cur)) / 2
	} else {
		return float64(cur)
	}
}

func findJudge(n int, trust [][]int) int {
	m := map[int]map[int]bool{}
	t := map[int]bool{}
	for _, arr := range trust {
		if m[arr[1]] == nil {
			m[arr[1]] = map[int]bool{}
		}
		m[arr[1]][arr[0]] = true
		t[arr[0]] = true
	}
	for k, arr := range m {
		if len(arr) == n-1 && !arr[k] && !t[k] {
			return k
		}
	}
	return -1
}

type WordDictionary struct {
	ch    [26]*WordDictionary
	isEnd bool
}

func (this *WordDictionary) AddWord(word string) {
	x := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if x.ch[idx] == nil {
			x.ch[idx] = &WordDictionary{}
		}
		x = x.ch[idx]
	}
	x.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	x := this
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for _, v := range x.ch {
				if v == nil {
					continue
				}
				if i == len(word)-1 {
					if v.isEnd == true {
						return true
					}
				} else if v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			idx := word[i] - 'a'
			if x.ch[idx] == nil {
				return false
			}
			x = x.ch[idx]
		}
	}
	return x.isEnd
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	big := make([]int, n)
	ans := []int{}
	pos := map[int]int{}
	stack := []int{-1}

	for i := n - 1; i >= 0; i-- {
		v := nums2[i]
		pos[v], big[i] = i, -1

		stack = append(stack, v)
		x := len(stack) - 2
		for ; x >= 0 && v > stack[x]; x-- {
			stack[x] = v
			stack = stack[:x+1]
		}
		if x >= 0 {
			big[i] = stack[x]
		}
		//fmt.Println(i, stack, x)
	}

	for i := 0; i < len(nums1); i++ {
		ans = append(ans, big[pos[nums1[i]]])
	}

	return ans
}

func countConsistentStrings(allowed string, words []string) int {
	m := map[int32]bool{}
	for _, v := range allowed {
		m[v] = true
	}
	ans := 0
	for _, v := range words {
		flag := true
		for _, x := range v {
			if !m[x] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	sum := make([]int, n)

	for i, v := range nums {
		if i == 0 {
			sum[i] = v
			continue
		}
		sum[i] = sum[i-1] + v
	}

	for i := 0; i < n; i++ {
		x := i * nums[i]
		if i > 0 {
			x -= sum[i-1]
		}

		x += sum[n-1] - sum[i]
		x -= (n - 1 - i) * nums[i]

		ans[i] = x
	}

	return ans
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	sum := make([]int, n)
	sb := 0
	for i := 0; i < n; i++ {
		sum[i] = aliceValues[i] + bobValues[i]
		sb += bobValues[i]
	}
	sort.Ints(sum)
	a := 0
	for i, _ := range sum {
		if i&1 == 0 {
			a += sum[n-1-i]
		}
	}

	ans := 0
	if a > sb {
		ans = 1
	} else if a < sb {
		ans = -1
	}
	return ans
}

func minDeletions(s string) int {
	cnt := make([]int, 26)
	for _, v := range s {
		cnt[v-'a']++
	}
	sort.Ints(cnt)
	ans := 0
	for i := 24; i >= 0; i-- {
		for cnt[i] > 0 && cnt[i] >= cnt[i+1] {
			cnt[i]--
			ans++
		}
	}
	return ans
}

func maxProfit(inventory []int, orders int) int {
	mod := int(1e9) + 7
	n := len(inventory)

	sort.Ints(inventory)
	sum := 0
	for _, v := range inventory {
		sum += v
	}

	left := sum - orders
	avg := left / n
	extra := left % n

	i := 0
	for inventory[i] <= avg {
		for ; i < n && inventory[i] <= avg; i++ {
			extra += avg - inventory[i]
		}
		avg += extra / (n - i)
		extra %= n - i
	}

	ans := 0
	for ; i < n; i++ {
		if extra == n-i {
			avg++
		}
		ans += (inventory[i] - avg) * (inventory[i] + avg + 1) / 2
		ans %= mod
	}
	return ans
}

func createSortedArray_(instructions []int) int {
	M := int(1e9) + 7
	n := len(instructions)
	//sort.Ints(instructions)
	//fmt.Println(instructions)

	arr := []int{}
	ans := 0

	for i := 0; i < n; i++ {
		l := sort.SearchInts(arr, instructions[i])
		r := sort.SearchInts(arr, instructions[i]+1)
		ans += min_t(l, len(arr)-r)
		ans %= M
		x := append([]int{}, arr[l:]...)
		arr = append(arr[:l], instructions[i])
		arr = append(arr, x...)
		fmt.Println(i, instructions[i], l, r, arr)
	}

	return ans
}
func min_t(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func removeInvalidParentheses(s string) []string {
	arr := []byte(s)
	m := map[int][]string{}
	backtrack(arr, 0, 0, m, []byte{})
	for i := len(arr); i >= 0; i-- {
		if m[i] != nil {
			hash := map[string]bool{}
			for _, v := range m[i] {
				hash[v] = true
			}
			ans := []string{}
			for k, _ := range hash {
				ans = append(ans, k)
			}
			return ans
		}
	}
	return []string{""}
}

func backtrack(arr []byte, i, l int, m map[int][]string, cur []byte) {
	//fmt.Println(i, l, cur)
	if i == len(arr) {
		if valid(cur) {
			if m[len(cur)] == nil {
				m[len(cur)] = []string{}
			}
			m[len(cur)] = append(m[len(cur)], string(cur))
		}
		return
	}
	if arr[i] >= 'a' && arr[i] <= 'z' {
		backtrack(arr, i+1, l, m, append(cur, arr[i]))
	} else if arr[i] == '(' {
		backtrack(arr, i+1, l, m, append([]byte{}, cur...))
		if l+1 <= len(arr)-i-1 {
			backtrack(arr, i+1, l+1, m, append(cur, arr[i]))
		}
	} else {
		backtrack(arr, i+1, l, m, append([]byte{}, cur...))
		if l > 0 {
			backtrack(arr, i+1, l-1, m, append(cur, arr[i]))
		}
	}
}

func valid(cur []byte) bool {
	left := 0
	for _, v := range cur {
		if v == '(' {
			left++
		} else if v == ')' {
			if left == 0 {
				return false
			}
			left--
		}
	}
	return left == 0
}

func reorderedPowerOf2(n int) bool {
	m := map[string]bool{}
	for i := 0; i <= 32; i++ {
		m[f(1<<i)] = true
	}

	if m[f(n)] {
		return true
	}
	return false
}

func f(n int) string {
	x := []int{}
	for n > 0 {
		x = append(x, n%10)
		n /= 10
	}
	sort.Ints(x)
	ans := []byte{}
	for _, v := range x {
		ans = append(ans, byte(v))
	}
	return string(ans)
}

// 树状数组
func createSortedArray(instructions []int) (cnt int) {
	mod := 1000000007
	tree := make([]int, 100001)
	var sum func(i int) int
	sum = func(i int) (x int) {
		for j := i; j > 0; j -= (j & -j) {
			x += tree[j]
		}
		return
	}
	for _, i := range instructions {
		for j := i; j < len(tree); j += (j & -j) {
			tree[j]++
			fmt.Println(j, tree[j])
		}
		cnt = (cnt + min(sum(i-1), sum(100000)-sum(i))) % mod
	}
	fmt.Println(sum(100000))
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPathCrossing(path string) bool {
	m := map[int]map[int]bool{0: {0: true}}
	cur := []int{0, 0}
	mv := map[byte][]int{'N': []int{0, 1}, 'E': []int{1, 0}, 'S': []int{0, -1}, 'W': []int{-1, 0}}
	for _, v := range path {
		idx := mv[byte(v)]
		cur[0] += idx[0]
		cur[1] += idx[1]
		if m[cur[0]] == nil {
			m[cur[0]] = map[int]bool{}
		}
		if m[cur[0]][cur[1]] {
			return true
		}
		m[cur[0]][cur[1]] = true
	}
	return false
}

func canArrange(arr []int, k int) bool {
	m := map[int]int{}
	for _, v := range arr {
		x := v % k
		if x < 0 {
			x += k
		}
		m[x]++
	}
	for i, v := range m {
		if i == 0 {
			if v&1 == 1 {
				return false
			}
		} else {
			if v != m[k-i] {
				return false
			}
		}
	}
	return true
}

func numSubseq(nums []int, target int) int {
	M := int(1e9) + 7
	ans := 0
	n := len(nums)
	sort.Ints(nums)
	i, j := 0, n-1
	for i < j {
		if nums[i]+nums[j] <= target {
			ans += q_p(j - i)
			ans %= M
			i++
		} else {
			j--
		}
	}
	if 2*nums[i] <= target {
		ans++
	}
	return ans % M
}

func q_p(a int) int {
	M := int(1e9) + 7
	ans := 1
	p := 2
	for a > 0 {
		if a&1 == 1 {
			ans *= p
			ans %= M
		}
		a /= 2
		p *= p
		p %= M
	}
	return ans
}

func modifyString(s string) string {
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] == '?' {
			arr[i] = 'a'
			for i > 0 && arr[i] == arr[i-1] || i <= len(arr)-2 && arr[i] == arr[i+1] {
				arr[i]++
			}
		}
	}
	return string(arr)
}

func numTriplets(nums1 []int, nums2 []int) int {
	sq1, sq2, p1, p2 := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	for _, v := range nums1 {
		sq1[v*v]++
	}
	for _, v := range nums2 {
		sq2[v*v]++
	}
	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			p1[nums1[i]*nums1[j]]++
		}
	}
	for i := 0; i < len(nums2)-1; i++ {
		for j := i + 1; j < len(nums2); j++ {
			p2[nums2[i]*nums2[j]]++
		}
	}

	ans := 0
	for k, v := range sq1 {
		ans += v * p2[k]
	}
	for k, v := range sq2 {
		ans += v * p1[k]
	}

	return ans
}

func countSubstrings(s string, t string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				continue
			}
			l, r := 1, 1
			for i-l >= 0 && j-l >= 0 && s[i-l] == t[j-l] {
				l++
			}
			for i+r < len(s) && j+r < len(t) && s[i+r] == t[j+r] {
				r++
			}
			ans += l * r
		}
	}
	return ans
}

func countSubstrings_(s string, t string) int {
	ans := 0
	a := count(s)
	b := count(t)
	for v, cnt := range a {
		for i := 0; i < len(v); i++ {
			for j := 'a'; j <= 'z'; j++ {
				if j == int32(v[i]) {
					continue
				}
				x := []byte(v)
				x[i] = byte(j)
				ans += cnt * b[string(x)]
			}
		}
	}
	return ans
}

func count(s string) map[string]int {
	cnt := map[string]int{}
	//sum := map[string]int{}
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		cur := []byte{}
		for j := i; j < len(arr); j++ {
			cur = append(cur, arr[j])
			cnt[string(cur)]++
			//sum[string(cur[:len(cur)-1])]++
		}
	}
	return cnt
}

func main() {
	fmt.Println(numOfArrays(2, 3, 1))
	//fmt.Println(numOfArrays(50,100,25))
	fmt.Println(numOfArrays(9, 1, 1))
}

func numWays(words []string, target string) int {
	M := int(1e9) + 7
	ans := 0
	cnt := make([]map[int]int, 26) // 记录每个字符每个index出现次数

	mx, n := 0, len(target)
	for i, _ := range cnt {
		cnt[i] = map[int]int{}
	}
	for _, s := range words {
		if len(s) > mx {
			mx = len(s)
		}
		for i, v := range s {
			cnt[v-'a'][i]++
		}
	}

	dp := make([][]int, len(target)) // dp[i][k] target第i个字符选到了k位置
	for i := 0; i < n; i++ {
		dp[i] = make([]int, mx)
		st := target[i] - 'a'
		tmp := 0
		for k := i; k < mx-(n-1-i); k++ {
			if i == 0 {
				tmp = 1
			} else {
				tmp += dp[i-1][k-1]
			}
			if cnt[st][k] > 0 {
				// dp[i][k] = dp[i-1][{k所有可能的位置累加}]
				dp[i][k] += tmp * cnt[st][k]
				dp[i][k] %= M
			}
		}
	}

	for _, v := range dp[n-1] {
		ans += v
		ans %= M
	}
	return ans
}

func numOfArrays(n int, m int, k int) int {
	if k > m {
		return 0
	}

	M := int(1e9) + 7
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	var dfs func(i, j, x, cur int)
	dfs = func(i0, j0, x0, cur int) { // i当前位置  j当前值 x剩余coast
		if x0 == 0 {
			dp[i0][j0][x0] = cur * qckP(j0+1, n-i0) % M
			return
		}

		t := 0
		for i := i0; i < n-1; i++ {
			for j := j0; j < m; j++ {
				if dp[i+1][j][x0-1] == 0 {
					dfs(i+1, j, x0-1, cur*qckP(j0+1, i-i0))
				}
				t += cur * dp[i+1][j][x0-1] % M
			}
		}
		dp[i0][j0][x0] = t % M
	}

	ans := 0
	for j := 0; j < m; j++ {
		dfs(1, j, k-1, 1)
		ans += dp[0][j][k]
	}
	fmt.Println(dp)

	return ans
}

func qckP(a, n int) int {
	M := int(1e9) + 7
	ans := 1
	for n > 0 {
		if n&2 == 1 {
			ans *= a
			ans %= M
		}
		n /= 2
		a *= a
		a %= M
	}
	return ans % M
}
