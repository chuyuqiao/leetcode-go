package problem0020

// isValid 判断字符串是否是有效的括号组成
func isValid(s string) bool {
	stack := make([]byte, len(s))
	top := 0
	for i, _ := range s {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}

		}
	}
	return top == 0
}
