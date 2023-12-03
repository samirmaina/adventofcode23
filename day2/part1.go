package main

func part1(bag set, games []game) int {
	var sum int
	for _, game := range games {
		if validateGame(bag, game) {
			sum += game.id
		}
	}
	return sum
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
