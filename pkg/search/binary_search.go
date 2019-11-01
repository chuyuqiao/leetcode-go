package search

// BinarySearch 二分查找算法
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

func bsearch(a []int, n, val int) int {
	return bsearchInternally(a, 0, n-1, val)
}

func bsearchInternally(a []int, low, high, val int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if a[mid] == val {
		return mid
	}
	if a[mid] > val {
		return bsearchInternally(a, low, mid-1, val)
	}
	if a[mid] < val {
		return bsearchInternally(a, mid+1, high, val)
	}
	return -1
}
