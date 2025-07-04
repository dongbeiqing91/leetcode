package easy

// https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=dynamic-programming
// 70. 爬楼梯

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var res int
	a1 := 1
	a2 := 2
	for i := 3; i <= n; i++ {
		res = a1 + a2
		a1 = a2
		a2 = res
	}
	return res
}
