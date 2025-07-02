package hard

import (
	"log"
	"math"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 4. 寻找两个正序数组的中位数

// 我这个答案时间复杂度达不到要求。其实用 sort.Ints直接排序也是一样的，如findMedianSortedArrays2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return 0.0
	}
	if len(nums1) == 0 {
		nums1 = nums2
		nums2 = nil
	}
	appended := append(nums1, nums2...)
	if totalLen == 1 {
		return float64(appended[0])
	}
	if totalLen == 2 {
		return float64(appended[0]+appended[1]) / 2
	}
	list := convertArrayToList(nums1)
	first := list
	for _, num := range nums2 {
		for list != nil {
			cur := list.Val
			if num < cur {
				prev := new(MyListNode)
				prev.Val = num
				prev.Next = list
				first = prev
				list = first
				break
			}
			// num >= cur
			// 插入队尾
			if list.Next == nil {
				list.Next = new(MyListNode)
				list.Next.Val = num
				list = list.Next
				break
			} else {
				next := list.Next.Val
				// 插入中间
				if num <= next {
					tmp := list.Next
					list.Next = new(MyListNode)
					list.Next.Val = num
					list.Next.Next = tmp
					list = list.Next
					break
				} else {
					list = list.Next
				}
			}
		}
	}
	pos := 0
	list = first
	if totalLen%2 != 0 {
		pos = int(math.Floor(float64(totalLen) / 2))
		for i := 0; i <= pos-1; i++ {
			list = list.Next
		}
		return float64(list.Val)
	}
	pos = totalLen/2 - 1
	for i := 0; i <= pos-1; i++ {
		list = list.Next
	}
	val1 := list.Val
	val2 := list.Next.Val
	return float64(val1+val2) / 2
}

type MyListNode struct {
	Val  int
	Next *MyListNode
}

func convertArrayToList(nums []int) *MyListNode {
	list := new(MyListNode)
	first := list
	for index, num := range nums {
		list.Val = num
		if index != len(nums)-1 {
			list.Next = new(MyListNode)
			list = list.Next
		}
	}
	return first
}

// 时间复杂度达不到要求
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	totalLen := len(nums)
	sort.Ints(nums)
	pos := 0
	if totalLen%2 != 0 {
		pos = int(math.Floor(float64(totalLen) / 2))
		return float64(nums[pos])

	}
	pos = totalLen/2 - 1

	val1 := nums[pos]
	val2 := nums[pos+1]
	return float64(val1+val2) / 2
}

func TestResult2(t *testing.T) {
	res := findMedianSortedArrays([]int{}, []int{1, 2, 3, 4})
	log.Print(res)
	res2 := findMedianSortedArrays2([]int{}, []int{1, 2, 3, 4})
	log.Print(res2)
}
