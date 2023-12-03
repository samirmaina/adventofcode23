package main

import (
	"log"
	"strconv"
	"unicode"
)

const part1SampleInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

type p1 struct{}

func part1(rows []string) int {

	var sum int
	for _, s := range rows {
		number := p1{}.findFirst(s) + p1{}.findSecond(s)
		sum += number
	}

	return sum
}

func (p1) findFirst(s string) int {
	for _, r := range s {
		if unicode.IsDigit(r) {
			x, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			}
			return 10 * x
		}
	}
	return 0
}

func (p1) findSecond(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if unicode.IsDigit(r) {
			x, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			}
			return x
		}
	}
	return 0
}
