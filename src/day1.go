package src

/*--- Day 1: Not Quite Lisp ---*/

func Day1(s string) (int, int) {
	floor := 0
	basementchar := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			floor++
		} else if s[i] == ')' {
			floor--
		}

		if floor == -1 && basementchar == -1 {
			basementchar = i + 1
		}
	}
	return floor, basementchar
}
