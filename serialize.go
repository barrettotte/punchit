package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// create a deck of punchcards from a file
func punchFile(path string) {
	lines, err := readLines(path)
	check(err)

	var deck []*punchcard

	for _, line := range lines {
		card := punchcard{}
		card.src = line
		deck = append(deck, &card)
	}

	for _, card := range deck {
		fmt.Println("\n", card)
	}
}

// read file into lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
