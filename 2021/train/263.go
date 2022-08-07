package main

type Bank struct {
	mon []int64
}

func Constructor_(balance []int64) Bank {
	return Bank{mon: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > len(this.mon) || account2 > len(this.mon) {
		return false
	}
	if this.mon[account1-1] >= money {
		this.mon[account1-1] -= money
		this.mon[account2-1] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > len(this.mon) {
		return false
	}
	this.mon[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > len(this.mon) {
		return false
	}
	if this.mon[account-1] >= money {
		this.mon[account-1] -= money
		return true
	}
	return false
}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	m := map[int]int{}
	max := 0
	cur := []int{}
	for i := 1; i < 1<<n; i++ {
		cur = []int{}
		for j := 0; j < n; j++ {
			if i&1<<j > 0 {
				cur = append(cur, nums[j])
			}
		}
		x := 0
		for _, v := range cur {
			x |= v
		}
		if x > max {
			max = x
		}
		m[x]++
	}
	return m[max]
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	sub := make([][]int, n+1)
	vis := map[int]bool{}
	for _, e := range edges {
		sub[e[0]] = append(sub[e[0]], e[1])
		sub[e[1]] = append(sub[e[1]], e[0])
	}

	next := map[int]bool{}
	cur := map[int]bool{1: true}
	deep, flag := 0, true
	for flag {
		deep++
		next = map[int]bool{}
		for e, _ := range cur {
			for _, v := range sub[e] {
				if vis[v] {
					continue
				}
				vis[v] = true
				next[v] = true
				if v == n {
					flag = false
				}
			}
		}
		cur = next
	}

	flag = false
	for e, _ := range cur {
		if flag {
			break
		}
		for _, v := range sub[e] {
			if v == n {
				deep++
				flag = true
				break
			}
		}
	}
	if !flag {
		deep += 2
	}

	ans := 0
	for i := 0; i < deep-1; i++ {
		ans += time
		if (ans/change)%2 == 0 {
			ans += change - ans%change
		}
	}
	return ans + time
}
