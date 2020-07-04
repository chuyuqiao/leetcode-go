package problem0498

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	n, m := len(matrix), len(matrix[0])

	var result []int
	// 对角线
	var diagonalLine []int

	for i := 0; i < n+m-1; i++ {
		diagonalLine = diagonalLine[:0]
		// x
		var x int
		if i < m {
			x = 0
		} else {
			x = i - m + 1
		}
		// y
		var y int
		if i < m {
			y = i
		} else {
			y = m - 1
		}

		for x < n && y > -1 {
			diagonalLine = append(diagonalLine, matrix[x][y])
			x++
			y--
		}

		if i%2 == 0 {
			diagonalLine = reverse(diagonalLine)
		}

		result = append(result, diagonalLine...)

	}

	return result
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
