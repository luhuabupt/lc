package main

func main() {
	n := int(1e7)
	Pirme(n)
	EulerPrime(n)
}

// 埃氏素数筛
func Pirme(n int) []int {
	m := make([]bool, n+1)
	for i := 2; i <= n/2; i++ {
		for j := 2; i*j <= n; j++ {
			m[i*j] = true
		}
	}

	ans := []int{}
	for i := 2; i <= n; i++ {
		if !m[i] {
			ans = append(ans, i)
		}
	}

	return ans
}

// 欧拉素数筛
// 合数 = 最小质因数 * 最大因数（非自己）
// 保证每个合数，都是被 最大因数 筛掉
func EulerPrime(n int) []int {
	m := make([]bool, n+1)
	ans := []int{}
	for i := 2; i <= n; i++ {
		if !m[i] { // isPrime
			ans = append(ans, i)
		}
		for _, v := range ans {
			if i*v > n {
				break
			}
			m[i*v] = true
			if i%v == 0 {
				break
			}
		}
	}
	return ans
}
