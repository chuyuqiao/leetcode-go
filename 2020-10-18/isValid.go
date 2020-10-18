package _020_10_18

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for _, b := range []byte(s) {
		if len(stack) == 0 || stack[len(stack)-1] != pairs[b] {
			stack = append(stack, b)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
