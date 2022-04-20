package src

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

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
	fmt.Println(floor)
	fmt.Println(basementchar)
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

	fmt.Println(area)
	fmt.Println(ribbon)

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

	fmt.Println(len(r))
}

func TestDay4(t *testing.T) {
	fmt.Println(Day4("bgvyzdsv"))
}
