package _006_zigzag_conversion

func convert(s string, numRows int) string {
	if numRows == 0 {
		return s
	}

	n := min(numRows, len(s))
	z := make([]string, n)
	res := ""

	curRow, goingDown := 0, false
	for _, c := range s {
		z[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if !goingDown {
			curRow -= 1
		} else {
			curRow += 1
		}
	}

	for _, str := range z {
		res += str
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
