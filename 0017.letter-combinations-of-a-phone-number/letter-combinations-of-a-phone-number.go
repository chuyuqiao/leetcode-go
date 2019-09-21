package problem0017

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	//dummy ret
	ret := []string{""}
	for i, d := range digits {
		//temp variable
		var temp []string
		for _, r := range ret {
			for k, _ := range m[digits[i]] {
				temp = append(temp, r+m[byte(d)][k])
			}
		}
		ret = temp
	}
	return ret
}
