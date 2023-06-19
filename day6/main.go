package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filebuffer, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	file := string(filebuffer)
	fmt.Println(string(filebuffer))
	for x := 0; x < len(file)-4; x++ {
		if unique(file[x : x+4]) {
			fmt.Println(file[x : x+4])
			fmt.Printf("%d\n", x+4)
			break
		}
	}
	for x := 0; x < len(file)-14; x++ {
		if unique(file[x : x+14]) {
			fmt.Println(file[x : x+14])
			fmt.Printf("%d\n", x+14)
			break
		}
	}
}

// check if all characters are unique in a slice
func unique(s string) bool {
	seen := make(map[int]bool)
	for _, v := range s {
		vi := int(v)
		if seen[vi] {
			return false
		}
		seen[vi] = true
	}
	return true
}
