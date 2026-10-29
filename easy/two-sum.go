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
	hash := map[int]int{}
	for index, num := range nums {
		expected := target - num
		if expectedIndex, ok := hash[expected]; ok {
			return []int{index, expectedIndex}
		}
		hash[num] = index
	}
	return nil
}
