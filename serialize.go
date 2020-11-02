package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type punchError struct {
	line   int
	col    int
	detail string
}

func (e *punchError) Error() string {
	return fmt.Sprintf("Error occurred punching line %d column %d.\n%s", e.line, e.col, e.detail)
}

// create a deck of punchcards from a file
func punchFile(path string) ([]*punchcard, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}
	return punchLines(lines)
}

// create a deck of punchcards from a slice of strings
func punchLines(lines []string) ([]*punchcard, error) {
	var deck []*punchcard
	var line string

	// punch a card for each line
	for i := 0; i < len(lines); i++ {
		if i >= len(lines) {
			line = strings.Repeat(" ", cols)
		} else {
			line = fmt.Sprintf("%*s", -cols, lines[i]) // pad to max column size
			line = line[0:cols]                        // trim excess characters
		}
		card, err := punchLine(line)

		if err != nil {
			perr := err.(*punchError)
			perr.line = i + 1
			return deck, perr
		}
		deck = append(deck, card)
	}
	return deck, nil
}

// create a punchcard from a single line
func punchLine(line string) (*punchcard, error) {
	card := punchcard{}
	card.src = line

	for j, col := range line {
		punches, err := punchChar(string(col))

		if err != nil {
			perr := err.(*punchError)
			perr.col = j + 1
			return nil, perr
		}
		card.punches[j] = punches
	}
	return &card, nil
}

// punch a char to 12-bit IBM029 encoding
func punchChar(c string) (string, error) {
	if len(c) > 1 {
		return "", &punchError{1, 1, fmt.Sprintf("cannot punch multiple characters '%s'", c)}
	}
	if col, ok := ibm029[strings.ToUpper(c)]; ok {
		return col, nil
	}
	return "", &punchError{1, 1, fmt.Sprintf("encountered invalid character '%s'", c)}
}

// read file into lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
