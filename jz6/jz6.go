package main

import "fmt"

func main() {
	fmt.Println(minNumberInRotateArray([]int{3, 4, 5, 1, 2}))
}

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}
	min := rotateArray[0]
	for i := 0; i < len(rotateArray); i++ {
		if rotateArray[i] < min {
			min = rotateArray[i]
		}
	}
	return min
}
