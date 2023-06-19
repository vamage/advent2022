package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	lowerCase = 96
	upperCase = 38
)

func main() {
	data := loadFile("input.txt")
	totalPriority := findPriority(data)
	groupPriority := findGroupPriority(data)

	fmt.Printf("part one, %v\n", totalPriority)
	fmt.Printf("part two, %v\n", groupPriority)
}

func loadFile(filename string) []string {
	lines := []string{}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func findPriority(lines []string) int {
	i := 0
	for _, line := range lines {
		// find shared letter in each string

		split := []string{line[0 : len(line)/2], line[len(line)/2:]}
	out:
		for _, letter := range split[1] {
			for _, letter2 := range split[0] {
				if letter2 == letter {
					fmt.Printf("%c, %d \n", letter, lookupPriority(letter))
					i += lookupPriority(letter)
					break out
				}
			}
		}

	}

	return i
}

func lookupPriority(in int32) int {
	if in > lowerCase {
		return int(in - lowerCase)
	} else {
		return int(in - upperCase)
	}
}

func findGroupPriority(lines []string) int {
	i := 0

	for c := 0; c < len(lines); c += 3 {
	out:
		for _, letter := range lines[c] {
			for _, letter2 := range lines[c+1] {
				if letter == letter2 {
					for _, letter3 := range lines[c+2] {
						if letter == letter3 {
							i += lookupPriority(letter)
							break out
						}
					}
				}
			}
		}
	}

	return i
}
