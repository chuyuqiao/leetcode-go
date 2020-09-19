package __queens

import "fmt"

func solveNQueens(n int) {
	result := make([]int, n)
	calNQueens(0, n, result)
}

func calNQueens(row, n int, result []int) {
	if row == n {
		printNQueens(n, result)
		return
	}
	for column := 0; column < n; column++ {
		if isOkn(row, column, result) {
			result[row] = column
			calNQueens(row+1, n, result)
		}
	}
}

func isOkn(row, column int, result []int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i > 0; i-- {
		if result[i] == column {
			return false
		}
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
		}
		if rightup < 8 {
			if result[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

func printNQueens(n int, result []int) {
	for row := 0; row < n; row++ {
		for column := 0; column < n; column++ {
			if result[row] == column {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
