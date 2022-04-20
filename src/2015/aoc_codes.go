package src

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

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

/*--- Day 3: Perfectly Spherical Houses in a Vacuum ---*/

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

/*--- Day 4: The Ideal Stocking Stuffer ---*/

func Day4(secretKey string) (int, error) {
	result := 0
	for {
		r := strconv.FormatInt(int64(result), 10)
		str := secretKey + r
		md5Result := GetMD5Hash(str)
		if md5Result[0:6] == "000000" {
			break
		}
		result++
	}

	return result, nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

/*--- Day 5: Doesn't He Have Intern-Elves For This? ---*/

func checkCond1(s string) bool {

	counter := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			counter++
		}
	}

	if counter >= 3 {
		return true
	}

	return false
}

func checkCond2(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func checkCond3(s string) bool {
	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
		return false
	}
	return true
}

// func checkCond4(s string) bool {
// 	for i:=0;i<len(s);i++{

// 	}
// }

func Day5Part1(s string) bool {

	if checkCond1(s) && checkCond2(s) && checkCond3(s) {
		return true
	}
	return false
}
