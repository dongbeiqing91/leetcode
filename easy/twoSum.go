package easy

func twoSum(nums []int, target int) []int {
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
