package problem0017

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
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
	ret := []string{""}
	for _, d := range digits {
		var temp []string
		for _, r := range ret {
			for _, l := range m[byte(d)] {
				temp = append(temp, r+l)
			}
		}
		ret = temp
	}
	return ret
}
