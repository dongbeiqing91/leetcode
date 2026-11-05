package easy

// https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-i

func incremovableSubarrayCount(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for cutLen := 1; i+cutLen <= len(nums); cutLen++ {
			n1 := nums[0:i]
			n2 := nums[i+cutLen:]
			if isIncremental(n1, n2) {
				count++
			}
		}
	}
	return count
}

func isIncremental(n1, n2 []int) bool {
	if !isIncNum(n1) || !isIncNum(n2) {
		return false
	}
	if len(n1) == 0 || len(n2) == 0 {
		return true
	}
	if n1[len(n1)-1] < n2[0] {
		return true
	}
	return false
}

func isIncNum(num []int) bool {
	if len(num) == 0 {
		return true
	}
	for i := 0; i < len(num)-1; i++ {
		if num[i+1] <= num[i] {
			return false
		}
	}
	return true
}
