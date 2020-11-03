package punchit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PunchError represents where an error with punching occurred
type PunchError struct {
	line   int
	col    int
	detail string
}

func (e *PunchError) Error() string {
	return fmt.Sprintf("Error occurred punching line %d column %d.\n%s", e.line, e.col, e.detail)
}

// PunchFile punches a deck of punchcards from a file
func PunchFile(path string) ([]*Punchcard, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}
	return PunchLines(lines)
}

// PunchLines punches a deck of punchcards from a slice of strings
func PunchLines(lines []string) ([]*Punchcard, error) {
	var deck []*Punchcard
	var line string

	// punch a card for each line
	for i := 0; i < len(lines); i++ {
		if i >= len(lines) {
			line = strings.Repeat(" ", cols)
		} else {
			line = lines[i]
		}
		card, err := PunchLine(line)

		if err != nil {
			perr := err.(*PunchError)
			perr.line = i + 1
			return deck, perr
		}
		deck = append(deck, card)
	}
	return deck, nil
}

// PunchLine punches a punchcard from a single line
func PunchLine(s string) (*Punchcard, error) {
	line := fmt.Sprintf("%*s", -cols, s) // pad to max column size
	line = line[0:cols]                  // trim excess characters

	card := Punchcard{}
	card.src = line

	for j, col := range line {
		punches, err := PunchChar(string(col))

		if err != nil {
			perr := err.(*PunchError)
			perr.col = j + 1
			return nil, perr
		}
		card.punches[j] = punches
	}
	return &card, nil
}

// PunchChar punches a char to a 12-bit binary string with IBM029 encoding
func PunchChar(c string) (string, error) {
	if len(c) > 1 {
		return "", &PunchError{1, 1, fmt.Sprintf("cannot punch multiple characters '%s'", c)}
	}
	if col, ok := ibm029[strings.ToUpper(c)]; ok {
		return col, nil
	}
	return "", &PunchError{1, 1, fmt.Sprintf("encountered invalid character '%s'", c)}
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
