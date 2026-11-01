package easy

// https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=dynamic-programming
// 70. 爬楼梯

// c(1) = 1
// c(2) = 2
// c(n) = c(n-1) + c(n-2)
// leetcode里提交递归会超时, 要改成循环
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return n
	}

	prev, current := 1, 2
	for i := 3; i <= n; i++ {
		prev, current = current, prev+current
		/* or:
		next := prev + current
		prev = current
		current = next
		*/
	}
	return current
}
