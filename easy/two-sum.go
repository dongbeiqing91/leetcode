package easy

// https://leetcode.cn/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for index, num := range nums {
		x := target - num
		if xIndex, ok := m[x]; ok {
			return []int{index, xIndex}
		}
		m[num] = index // 这一行不能放到 for 循环的最前面, 避免找到自己, 比如nums: [3,2,4], target=6, 6-3=3 这种情况
	}
	return nil
}
