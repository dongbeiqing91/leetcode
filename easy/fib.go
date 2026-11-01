package easy

// 0、1、1、2、3、5、8、13、21、34、55、89、144
func fib(n int) int {
	if n <= 1 {
		return n
	}
	prev := 0
	curr := 1
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
