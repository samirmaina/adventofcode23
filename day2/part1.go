package main

import (
	"strconv"
	"strings"
)

type color string
type set map[color]int
type game struct {
	id   int
	sets []set
}

func part1(bag set, rows []string) int {
	var sum int
	for _, row := range rows {
		g := parseRow(row)
		if validateGame(bag, g) {
			sum += g.id
		}
	}

	return sum
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

func validateGame(bag set, g game) bool {
	// if any of the game sets has more of a specific color in the bag, bail
	for _, set := range g.sets {
		for c, count := range set {
			if count > bag[c] {
				return false
			}
		}
	}
	return true
}
