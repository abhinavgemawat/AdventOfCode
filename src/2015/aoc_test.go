package src

import (
	"bufio"
	"os"
	"testing"
)

func CmpInt(expected int, actual int) bool {
	if expected == actual {
		return true
	}
	return false
}
func TestDay1(t *testing.T) {
	file := "testdata/day1.txt"

	h, err := os.Open(file)
	if err != nil {
		t.Fatal("error while opening file", err)
	}
	defer h.Close()

	var floor, basementchar int
	scan := bufio.NewScanner(h)
	for scan.Scan() {
		floor, basementchar = Day1(scan.Text())
	}

	CmpInt(280, floor)
	CmpInt(1797, basementchar)
}

func TestDay2(t *testing.T) {
	file := "testdata/day2.txt"

	h, err := os.Open(file)
	if err != nil {
		t.Fatal("error while opening file", err)
	}
	defer h.Close()

	var tarea, trib []int
	scan := bufio.NewScanner(h)
	for scan.Scan() {
		area, rl, err := Day2(scan.Text())

		if err != nil {
			t.Fatal("could not read file", err)
		}
		tarea = append(tarea, area)
		trib = append(trib, rl)
	}

	area := 0
	ribbon := 0
	for i := 0; i < len(tarea); i++ {
		area += tarea[i]
		ribbon += trib[i]
	}

	CmpInt(1709812, area)
	CmpInt(3863354, ribbon)

}

func TestDay3(t *testing.T) {
	file := "testdata/day3.txt"

	h, err := os.Open(file)
	if err != nil {
		t.Fatal("error while opening file", err)
	}
	defer h.Close()

	scan := bufio.NewScanner(h)

	r := make(map[Coordinate]int)
	for scan.Scan() {
		r, err = Day3(scan.Text())
		if err != nil {
			t.Fatal("could not read data", err)
		}
	}

	CmpInt(2360, len(r))
}

func TestDay4(t *testing.T) {
	v, err := Day4("bgvyzdsv")
	if err != nil {
		t.Fatal("could not execute day 4")
	}
	CmpInt(1038736, v)
}

func TestDay5(t *testing.T) {
	file := "testdata/day5.txt"

	h, err := os.Open(file)
	if err != nil {
		t.Fatal("error while opening file", err)
	}
	defer h.Close()

	scan := bufio.NewScanner(h)

	counter := 0
	for scan.Scan() {
		r := Day5Part1(scan.Text())
		if r {
			counter++
		}
	}

	CmpInt(258, counter)
}
