package easy

// https://leetcode.cn/problems/move-zeroes

func moveZeroes(nums []int) {
	k := removeZero(nums)
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

func removeZero(nums []int) int {
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}
