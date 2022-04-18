package src

import (
	"fmt"
	"strconv"
	"strings"
)

/*--- Day 2: I Was Told There Would Be No Math ---*/

func Day2(s string) (int, int, error) {
	arr := strings.Split(s, "x")

	area := 0

	l, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, fmt.Errorf("could not convert: %s", err)
	}

	w, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, 0, fmt.Errorf("could not convert: %s", err)
	}

	h, err := strconv.Atoi(arr[2])
	if err != nil {
		return 0, 0, fmt.Errorf("could not convert: %s", err)
	}

	var a_sides []int
	a_sides = append(a_sides, l*w)
	a_sides = append(a_sides, w*h)
	a_sides = append(a_sides, l*h)

	area += 2 * (a_sides[0] + a_sides[1] + a_sides[2])

	min_side_area := 10000

	for i := 0; i < len(a_sides); i++ {
		if min_side_area > a_sides[i] {
			min_side_area = a_sides[i]
		}
	}

	area += min_side_area

	ribbonlength := 0
	var p_sides []int
	p_sides = append(p_sides, l+w)
	p_sides = append(p_sides, w+h)
	p_sides = append(p_sides, l+h)

	min_side_p := 10000
	for i := 0; i < len(p_sides); i++ {
		if min_side_p > p_sides[i] {
			min_side_p = p_sides[i]
		}
	}

	ribbonlength += (2 * min_side_p) + (l * w * h)

	return area, ribbonlength, nil
}
