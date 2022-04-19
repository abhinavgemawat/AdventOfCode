package src

type Coordinate struct {
	x int
	y int
}

func Day3(s string) (map[Coordinate]int, error) {
	route := make(map[Coordinate]int)
	var santaloc Coordinate
	var roboloc Coordinate

	for i := 0; i < len(s); i++ {
		if i%2 != 0 {
			if s[i] == '>' {
				roboloc.x++
			}
			if s[i] == '<' {
				roboloc.x--
			}
			if s[i] == '^' {
				roboloc.y++
			}
			if s[i] == 'v' {
				roboloc.y--
			}
			route[roboloc]++
		} else {
			if s[i] == '>' {
				santaloc.x++
			}
			if s[i] == '<' {
				santaloc.x--
			}
			if s[i] == '^' {
				santaloc.y++
			}
			if s[i] == 'v' {
				santaloc.y--
			}
			route[santaloc]++
		}

	}

	return route, nil
}
