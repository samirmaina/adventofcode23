package main

import (
	"advent-of-code-23/utils"
	"fmt"
)

func main() {
	//rows := strings.Split(part1SampleInput, "\n")
	//rows := strings.Split(part2SampleInput, "\n")
	rows := utils.LoadInputFile(utils.Filepath("day1/input.txt"))

	fmt.Println(fmt.Sprintf("The first total is: %d", part1(rows)))
	// 55488

	fmt.Println(fmt.Sprintf("The second total is: %d", part2(rows)))
	// 55614
}
