package easy

// https://leetcode.cn/problems/n-th-tribonacci-number/description/

func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	var res int
	a1 := 0
	a2 := 1
	a3 := 1
	for i := 3; i <= n; i++ {
		res = a1 + a2 + a3
		a1 = a2
		a2 = a3
		a3 = res
	}
	return res
}

func tribonacci2(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
