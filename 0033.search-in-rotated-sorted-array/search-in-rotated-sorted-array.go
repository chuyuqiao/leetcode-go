package problem0033

func search(nums []int, target int) int {
	rotated := indexOfMin(nums)
	size := len(nums)
	low, high := 0, size-1
	for low <= high {
		mid := (low + high) / 2
		rotatedMid := (rotated + mid) % size
		if nums[rotatedMid] == target {
			return rotatedMid
		} else if nums[rotatedMid] > target {
			high = rotatedMid - 1
		} else {
			low = rotatedMid + 1
		}
	}
	return -1
}

func indexOfMin(nums []int) int {
	size := len(nums)
	low, high := 0, size-1
	for low < high {
		mid := (low + high) / 2
		if nums[high] < nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
