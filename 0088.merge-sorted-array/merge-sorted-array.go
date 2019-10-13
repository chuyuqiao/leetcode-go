package problem0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)
	//双指针
	j, k := 0, 0

	for i, _ := range nums1 {
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		if temp[j] <= nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}
