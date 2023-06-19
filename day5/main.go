package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := loadFile("input.txt")
	containers, comands := breaktoParts(data)
	loadedContainers := loadContainers(containers)
	loadedContainers9001 := loadContainers(containers)
	movedContainers := moveContainters(loadedContainers, comands)
	movedContainers9001 := moveContainters9001(loadedContainers9001, comands)
	for _, line := range containers {
		log.Println(line)
	}
	fmt.Println("-----")
	for _, line := range comands {
		log.Println(line)
	}
	fmt.Printf("%v", loadedContainers)
	for _, line := range movedContainers {
		if len(line) > 0 {
			fmt.Printf("%c", line[len(line)-1])
		}
	}
	fmt.Println("")
	for _, line := range movedContainers9001 {
		if len(line) > 0 {
			fmt.Printf("%c", line[len(line)-1])
		}
	}
	fmt.Println("")
}

func moveContainters(containers [][]int, commands []string) [][]int {
	fmt.Printf("%+v\n", containers)
	var m, source, destination, a int
	for _, command := range commands {
		_, err := fmt.Sscanf(command, "move %d from  %d to %d", &m, &source, &destination)
		if err != nil {
			log.Fatalf("failed to load moves %v", err)
		}
		for x := m; x > 0; x-- {

			a, containers[source] = containers[source][len(containers[source])-1], containers[source][:len(containers[source])-1]
			containers[destination] = append(containers[destination], a)
		}
	}
	return containers
}
func moveContainters9001(containers [][]int, commands []string) [][]int {
	fmt.Println("---9001---")
	fmt.Printf("%+v\n", containers)
	var m, source, destination int
	var a []int
	for _, command := range commands {
		_, err := fmt.Sscanf(command, "move %d from  %d to %d", &m, &source, &destination)
		if err != nil {
			log.Fatalf("failed to load moves %v", err)
		}

		a, containers[source] = containers[source][len(containers[source])-m:], containers[source][:len(containers[source])-m]
		containers[destination] = append(containers[destination], a...)
	}
	return containers
}

func loadContainers(containers []string) [][]int {
	a := make([][]int, 10)
	for x := 0; x < len(containers)-1; x++ {
		for y := 1; ((y-1)*4)+1 < len(containers[x])-1; y++ {
			pos := ((y - 1) * 4) + 1
			if containers[x][pos] == ' ' {
				fmt.Printf(" \n-+--\n")
			} else {
				a[y] = append([]int{int(containers[x][pos])}, a[y]...)

			}
		}
	}
	return a
}

func breaktoParts(input []string) ([]string, []string) {
	var container []string
	var command []string
	lines := 0
	for _, line := range input {
		if line == "" {
			lines = 1
		} else if lines == 0 {
			container = append(container, line)
		} else {
			command = append(command, line)
		}
	}

	return container, command
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
