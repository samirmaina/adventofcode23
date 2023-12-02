package main

import (
	"advent-of-code-23/utils"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	part1()
}

func part1() {
	const sampleInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	//rows := strings.Split(sampleInput, "\n")
	rows := utils.LoadInputFile(utils.Filepath("day1/input.txt"))

	var sum int
	for _, s := range rows {
		x, err := strconv.Atoi(string(findFirst(s)))
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(string(findSecond(s)))
		if err != nil {
			log.Fatal(err)
		}

		sum += (10 * x) + y

	}

	fmt.Println(fmt.Sprintf("The total is: %d", sum))
}

func findFirst(s string) rune {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return r
		}
	}
	return 0
}

func findSecond(s string) rune {
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if unicode.IsDigit(r) {
			return r
		}
	}
	return 0
}
