package main

import (
	"advent-of-code-23/utils"
	"fmt"
)

func main() {
	rows := utils.LoadInputFile(utils.Filepath("day2/input.txt"))
	fmt.Println(fmt.Sprintf("The first total is: %d", part1(set{
		"red":   12,
		"green": 13,
		"blue":  14,
	}, rows)))
}
