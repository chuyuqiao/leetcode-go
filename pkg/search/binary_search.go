package search

//BinarySearch 二分查找算法
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		meidan := (low + high) / 2
		if target < nums[meidan] {
			high = meidan - 1
		} else if target > nums[meidan] {
			low = meidan + 1
		} else {
			return meidan
		}

		/*switch {
		case nums[meidan] > target:
			high = meidan - 1
		case nums[meidan] < target:
			low = meidan + 1
		default:
			return meidan
		}*/
	}
	return -1
}
