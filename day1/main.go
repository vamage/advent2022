package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := loadFile("input.txt")
	elves := calories(lines)
	getlargestn(elves, 1)
	getlargestn(elves, 3)
}

func loadFile(filename string) []string {
	lines := make([]string, 0)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)

		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func calories(input []string) []int {
	elves := make([]int, 0)
	cal := 0
	c := 0
	for _, element := range input {
		if element == "" {

			fmt.Printf("Elf %d , calories %d \n", c+1, cal)
			c++
			elves = append(elves, cal)
			cal = 0
		} else {
			i, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}

			cal += i

		}
	}
	fmt.Printf("Elf %d , calories %d \n", c+1, cal)
	c++
	elves = append(elves, cal)
	cal = 0
	return elves
}

func getlargestn(input []int, n int) {
	max := make([]int, n, n)
	for _, element := range input {
		for k, v := range max {
			if element > v {
				tmp := v
				max[k] = element
				element = tmp

			}
		}
	}

	m := 0
	for _, v := range max {
		m += v
	}
	fmt.Printf("largest n elves %d\n", m)
}
