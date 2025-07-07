package easy

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
