package easy

// https://leetcode.cn/problems/sqrtx/description/

func mySqrt(x int) int {
	left := 0
	right := x
	ans := -1
	for left <= right {
		mid := (right-left)/2 + left
		sqr := mid * mid
		if sqr <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func mySqrt2(x int) int {
	for i := 0; i <= x; i++ {
		sqr := i * i
		if sqr == x {
			return i
		}
		sqr2 := (i + 1) * (i + 1)
		if sqr < x && sqr2 > x {
			return i
		}
	}
	return -1
}
