package _008_string_to_integer_atoi

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	return convert(clean(s))
}

func convert(sign int, absStr string) int {
	absNum := 0
	for _, b := range absStr {
		absNum = absNum*10 + int(b-'0')
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}

func clean(s string) (sign int, abs string) {
	// 去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	// 判断首字符
	fmt.Printf("%T", s[0])
	switch s[0] {
	// 有效
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
		//正号
	case '+':
		sign, abs = 1, s[1:]
		//负号
	case '-':
		sign, abs = -1, s[1:]
		//无效直接返回
	default:
		abs = ""
		return
	}
	for i, b := range abs {
		if b < '0' || b > '9' {
			abs = abs[:i]
			break
		}
	}
	return
}

func myAtoi2(s string) int {
	am := new()
	for _, r := range s {
		am.get(r)
	}
	return am.sign * am.ans
}

const (
	start int = iota
	signed
	in_number
	end
)

var stateMap = map[int][]int{
	start: {
		start, signed, in_number, end,
	},
	signed: {
		end, end, in_number, end,
	},
	in_number: {
		end, end, in_number, end,
	},
	end: {
		end, end, end, end,
	},
}

type automaton struct {
	state int
	sign  int
	ans   int
}

func new() *automaton {
	return &automaton{
		state: start,
		sign:  1,
		ans:   0,
	}
}

func (am *automaton) get(r rune) {
	am.state = stateMap[am.state][getCol(r)]
	switch am.state {
	case in_number:
		am.ans = am.ans*10 + int(r-'0')
		if am.sign == 1 && am.ans > math.MaxInt32 {
			am.ans = math.MaxInt32
		}
		if am.sign == -1 && am.ans*am.sign < math.MinInt32 {
			am.ans = -math.MinInt32
		}
	case signed:
		if '+' == r {
			am.sign = 1
		} else {
			am.sign = -1
		}
	}
}

func getCol(r rune) int {
	switch r {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	default:
		return 3
	}
}
