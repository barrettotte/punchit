package main

import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// test punch deck from file
	deck, err := punchFile("examples/hello.rpg")
	check(err)
	for _, card := range deck {
		fmt.Println(card)
	}

	// test punch single char
	col, err := punchChar("A")
	check(err)
	fmt.Println(col)

	// test punch card from string
	card, err := punchLine("Hello world")
	check(err)
	fmt.Println(card)

	// test punch deck from slice of strings
	lines := []string{"first card", "second card", "third card"}
	deck, err = punchLines(lines)
	for _, card := range deck {
		fmt.Println(card)
	}
}
