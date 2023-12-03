package main

import (
	"regexp"
)

const part2SampleInput = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
1eightwo`

func part2(rows []string) int {
	fwd := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[0-9])`)
	bwd := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9])`)
	var sum int
	for _, row := range rows {
		first := numMap[fwd.FindString(row)]
		second := numMap[reverse(bwd.FindString(reverse(row)))]
		number := (first * 10) + second
		//fmt.Println(fmt.Sprintf("%s %d", row, number))
		sum += number
	}
	return sum
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,

	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}
