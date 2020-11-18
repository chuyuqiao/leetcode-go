package sortx

func quickSort(nums []int) []int {
	quickSortN(nums, 0, len(nums)-1)
	return nums
}

func quickSortN(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(nums, p, r)
	quickSortN(nums, p, q-1)
	quickSortN(nums, q+1, r)
}

func partition(nums []int, p, r int) int {
	pivot := nums[r]
	i := p
	for j := p; j < r; j++ {
		if nums[j] < pivot {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
		}
	}
	nums[r] = nums[i]
	nums[i] = pivot
	return i
}
