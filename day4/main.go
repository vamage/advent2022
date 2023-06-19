package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := loadFile("input.txt")
	fullSubsets := findSubsets(data)
	partialSubsets := findSubset2(data)
	fmt.Printf("part 1: %d\n", fullSubsets)
	fmt.Printf("part 2: %d\n", partialSubsets)
}
func findSubsets(data []string) int {
	var a, b, c, d int
	count := 0
	for _, v := range data {
		_, err := fmt.Sscanf(v, "%d-%d,%d-%d", &a, &b, &c, &d)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d-%d,%d-%d\n", a, b, c, d)
		if (a <= c && b >= d) || (c <= a && d >= b) {
			fmt.Printf("#%d-%d,%d-%d\n", a, b, c, d)
			count++
		}
	}

	return count
}

func findSubset2(data []string) int {
	var a, b, c, d int
	count := 0
	for _, v := range data {
		_, err := fmt.Sscanf(v, "%d-%d,%d-%d", &a, &b, &c, &d)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d-%d,%d-%d\n", a, b, c, d)
		if a <= d && b >= c {
			fmt.Printf("#%d-%d,%d-%d\n", a, b, c, d)
			count++
		}
	}

	return count
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
