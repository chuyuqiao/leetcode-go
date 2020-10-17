package _020_10_16

func bubbleSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	for i := 0; i < n; i++ {
		flag := false
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

func insertionSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	for i := 1; i < n; i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
	return nums
}

func mergeSort(nums []int) []int {
	mergeSortN(nums, 0, len(nums)-1)
	return nums
}

func mergeSortN(nums []int, p, r int) []int {
	if p >= r {
		return nums
	}
	q := (p + r) / 2
	return merge(nums, mergeSortN(nums, p, q), mergeSortN(nums, q+1, r))
}

func merge(nums []int, p []int, r []int) []int {
	return nums
}

func quickSort(nums []int) []int {
	quickSortN(nums, 0, len(nums)-1)
	return nums
}

func quickSortN(nums []int, p int, r int) {
	if p >= r {
		return
	}
	q := partition(nums, p, r)
	quickSortN(nums, p, q-1)
	quickSortN(nums, q+1, r)
}

func partition(nums []int, p int, r int) int {
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
