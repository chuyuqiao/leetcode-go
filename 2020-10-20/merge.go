package _020_10_20

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0
	tmp := make([]int, m+n)
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else {
			tmp[k] = nums2[j]
			j++
		}
		k++
	}

	if i < m {
		for i < m {
			tmp[k] = nums1[i]
			i++
			k++
		}
	}

	if j < n {
		for j < n {
			tmp[k] = nums2[j]
			j++
			k++
		}
	}

	for i := 0; i < k; i++ {
		nums1[i] = tmp[i]
	}
}
