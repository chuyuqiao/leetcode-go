package problem0027

import "testing"

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	length := removeElement(nums, 3)
	t.Log(length)
	t.Log(nums[:length])
}
