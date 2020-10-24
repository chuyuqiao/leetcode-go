package _020_10_24

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	size := m + n
	nums3 := make([]int, size)
	merge(nums1, nums2, &nums3, m, n)
	if size%2 == 1 {
		return float64(nums3[size/2])
	} else {
		return float64(nums3[size/2-1]+nums3[size/2]) / 2
	}
}

func merge(nums1, nums2 []int, nums3 *[]int, m, n int) {
	k := 0
	if m == 0 {
		*nums3 = nums2
		return
	}
	if n == 0 {
		*nums3 = nums1
		return
	}
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			(*nums3)[k] = nums1[i]
			i++
		} else {
			(*nums3)[k] = nums2[j]
			j++
		}
		k++
	}

	if i < m {
		for i < m {
			(*nums3)[k] = nums1[i]
			i++
			k++
		}
	}

	if j < n {
		for j < n {
			(*nums3)[k] = nums2[j]
			j++
			k++
		}
	}
}
