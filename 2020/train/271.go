package main

func subArrayRanges(nums []int) int64 {
	n := len(nums)

	ans := int64(0)
	for i := 0; i < n; i++ {
		max, min := nums[i], nums[i]
		for j := i; j < n; j++ {
			if nums[i] > max {
				max = nums[i]
			}
			if nums[i] < min {
				min = nums[i]
			}
			ans += int64(max - min)
		}
	}
	return ans
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	n := len(plants)
	l, r := capacityA, capacityB
	ans := 0
	for i := 0; i < n/2; i++ {
		if l < plants[i] {
			ans++
			l = capacityA
		}
		if r < plants[n-1-i] {
			ans++
			r = capacityB
		}
	}
	if n%2 == 1 {
		if l < plants[n/2+1] && r < plants[n/2+1] {
			ans++
		}
	}
	return ans
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	m := map[int]int{}
	for _, v := range fruits {
		m[v[0]] = v[1]
	}

	dl, dr := []int{0}, []int{0}
	for i := 1; i <= k; i++ {
		dr = append(dr, dr[i-1]+m[startPos+i])
	}
	for i := 1; i <= k; i++ {
		dl = append(dl, dl[i-1]+m[startPos-i])
	}

	ans := 0
	for i := 0; i <= k; i++ {
		xl, xr := dl[i], dr[i]
		if k-2*i >= 0 {
			xl += dr[k-2*i]
			xr += dl[k-2*i]
		}
		if xl > ans {
			ans = xl
		}
		if xr > ans {
			ans = xr
		}
	}

	return ans + m[startPos]
}
