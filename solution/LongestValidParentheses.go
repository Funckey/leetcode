package solution

func longestValidParentheses(s string) int {
	l := len(s)
	records := make([]int, l)
	max := 0
	for i := 1; i < l; i++ {
		match := 0
		if s[i] == ')' {
			if s[i] == ')' {
				if s[i-1] == '(' {
					records[i] = 2
					match = 1
				}
				if records[i-1] > 0 && i - records[i-1] > 0 && s[i - records[i-1] - 1] == '(' {
					records[i] = records[i - 1] + 2
					match = 1
				}
				if match == 1 {
					if i - records[i] >=  0 && records[i - records[i]] > 0 {
						records[i] += records[i - records[i]]
					}
					if max < records[i] {
						max = records[i]
					}
				}
			}
		}
	}
	return max
}


