package main

import (
	"fmt"
	"github.com/ohoherror/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("C:\\Users\\User\\go\\bin\\votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
