package easy

// https://leetcode.cn/problems/remove-element/

func removeElement(nums []int, val int) int {
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}
