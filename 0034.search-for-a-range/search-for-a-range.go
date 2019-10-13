package problem0034

import "leetcode-go/pkg/search"

/**
给一个正序的数组，找出指定数字的开始和结束位置，不存在则[-1,-1]
*/
func searchRange(nums []int, target int) []int {

	index := search.BinarySearch(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	first := index
	for {
		f := search.BinarySearch(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	last := index
	for {
		l := search.BinarySearch(nums[last+1:], target)
		if l == -1 {
			break
		}
		last += l + 1
	}

	return []int{first, last}
}
