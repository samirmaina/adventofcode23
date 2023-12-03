package main

func part2(games []game) int {
	var powers []int

	// for each game
	// for each set
	// find the largest group of each color
	for _, g := range games {
		ideal := set{}
		for _, s := range g.sets {
			for color, i := range s {
				if current, ok := ideal[color]; !ok {
					ideal[color] = i
				} else if s[color] > current {
					ideal[color] = i
				}
			}
		}
		powers = append(powers, ideal.power())
	}
	var sum int
	for _, power := range powers {
		sum += power
	}
	return sum
}

func (s set) power() int {
	p := 1
	for _, i := range s {
		p *= i
	}
	return p
}
