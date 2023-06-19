package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type strat struct {
	you      int
	opponent int
}

func main() {
	lines := loadFile("input.txt")
	score(lines)
	scorev2(lines)
}

func loadFile(filename string) []strat {
	lines := make([]strat, 0)
	y := 0
	o := 0
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Sscanf(text, "%c %c", &o, &y)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, strat{you: y, opponent: o})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func score(input []strat) {
	sum := 0
	for _, element := range input {
		round := 0
		s := element.opponent - (element.you - 23)
		// fmt.Printf("%d %d, %d \n", element.opponent, (element.you), s)
		if element.you == 0 {
			// tie elements are equal
			round = 3
		} else if s == -1 || s == 2 {
			// win you are one off from your oppenent and wrap around
			round = 6
		}
		round = round + element.you - 87
		// fmt.Printf("round %d\n", round)
		sum += round
	}

	fmt.Printf("sum %d \n", sum)
}

func scorev2(input []strat) {
	sum := 0
	for _, element := range input {
		round := 0
		s := element.opponent - 64
		// fmt.Printf("%d %d, %d \n", element.opponent, (element.you ), s)
		if element.you == 89 {
			// tie Y= 89
			round = 3
		} else if element.you == 90 {
			// win Z =6
			round = 6
			s++
			if s > 3 {
				s = 1
			}
		} else {
			s--
			if s < 1 {
				s = 3
			}

		}
		round = round + s
		// fmt.Printf("round %d\n", round)
		sum += round
	}

	fmt.Printf("sum %d \n", sum)
}
