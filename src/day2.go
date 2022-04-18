package src

import (
	"fmt"
	"strconv"
	"strings"
)

/*--- Day 2: I Was Told There Would Be No Math ---*/

func Day2(s string) (int, int, error) {
	arr := strings.Split(s, "x")

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
	a_sides = append(a_sides, l)
	a_sides = append(a_sides, w)
	a_sides = append(a_sides, h)

	first := 10000
	second := 10000

	for i := 0; i < len(a_sides); i++ {
		if a_sides[i] < first {
			second = first
			first = a_sides[i]
		} else if a_sides[i] < second && a_sides[i] != first {
			second = a_sides[i]
		}
	}

	fmt.Println(first)
	fmt.Println(second)

	area := (2 * l * w) + (2 * w * h) + (2 * l * h) + first*second
	ribbonlength := 2*(first+second) + l*w*h

	return area, ribbonlength, nil
}
