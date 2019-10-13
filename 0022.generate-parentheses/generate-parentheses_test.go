package problem0022

import "testing"

func Test_generateParenthesis(t *testing.T) {
	parenthesis := generateParenthesis(3)
	for _, parenth := range parenthesis {
		t.Log(parenth)
	}

}
