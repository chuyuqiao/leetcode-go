package main

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace(s string) string {
	// write code here
	// return strings.ReplaceAll(s, " ", "%20")
	arr, size := make([]rune, len(s)*3), 0
	for _, c := range s {
		if c == ' ' {
			arr[size] = '%'
			size++
			arr[size] = '2'
			size++
			arr[size] = '0'
		} else {
			arr[size] = c
		}
		size++
	}
	return string(arr[:size])
}

func main() {
	rs := replaceSpace("We Are Happy")
	fmt.Println(rs)
}
