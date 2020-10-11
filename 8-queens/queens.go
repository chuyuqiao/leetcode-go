package __queens

import "fmt"

var result = [8]int{}

func cal8queens(row int) {
	if row == 8 {
		printQueens(result)
		return
	}
	for column := 0; column < 8; column++ {
		if isOk(row, column) {
			result[row] = column
			cal8queens(row + 1)
		}
	}
}

func isOk(row int, column int) bool {
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

func printQueens(result [8]int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
