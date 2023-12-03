package main

import (
	"advent-of-code-23/utils"
	"fmt"
	"strconv"
	"strings"
)

type color string
type set map[color]int
type game struct {
	id   int
	sets []set
}

func main() {
	//rows := utils.LoadInputFile(utils.Filepath("day2/example.txt"))
	rows := utils.LoadInputFile(utils.Filepath("day2/input.txt"))
	var games []game
	for _, row := range rows {
		games = append(games, parseRow(row))
	}

	fmt.Println(fmt.Sprintf("The first total is: %d", part1(
		set{"red": 12, "green": 13, "blue": 14},
		games,
	)))

	fmt.Println(fmt.Sprintf("The second total is: %d", part2(games)))
}

// Game 1: 9 red, 2 green, 13 blue; 10 blue, 2 green, 13 red; 8 blue, 3 red, 6 green; 5 green, 2 red, 1 blue
func parseRow(row string) game {
	g := game{}

	parts := strings.Split(row, ":")
	id, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(parts[0], "Game", "", 1)))
	g.id = id

	for _, setString := range strings.Split(parts[1], ";") {
		set := set{}
		for _, s := range strings.Split(setString, ",") {
			ps := strings.Split(strings.TrimSpace(s), " ")
			num, _ := strconv.Atoi(ps[0])
			c := color(ps[1])
			set[c] = num // <- could validate against bag right away
		}
		g.sets = append(g.sets, set)
	}

	return g
}
