package easy

import "sort"

// https://leetcode.cn/problems/maximum-units-on-a-truck/description/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var count int
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	for _, b := range boxTypes {
		if b[0] >= truckSize {
			count += truckSize * b[1]
			break
		}
		count += b[0] * b[1]
		truckSize -= b[0]
	}
	return count
}
