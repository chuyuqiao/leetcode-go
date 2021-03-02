package main

import "fmt"

var stack1 []int
var stack2 []int

func main() {
	Push(1)
	Push(2)
	Push(3)
	fmt.Println(Pop())
	fmt.Println(Pop())
	Push(4)
	Push(5)
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	Push(1)
	Push(2)
	fmt.Println(Pop())
	fmt.Println(Pop())
}

func Push(node int) {
	if len(stack2) > 0 {
		for i := len(stack2) - 1; i >= 0; i-- {
			stack1 = append(stack1, stack2[i])
		}
		stack2 = nil
	}
	stack1 = append(stack1, node)
}

func Pop() int {
	s1l, s2l := len(stack1), len(stack2)
	if s1l == 0 && s2l == 0 {
		return -1
	}
	if s1l > 0 {
		for i := len(stack1) - 1; i > 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		res := stack1[0]
		stack1 = nil
		return res
	}
	if s2l > 0 {
		res := stack2[s2l-1]
		stack2 = stack2[:s2l-1]
		return res
	}
	return -1
}
