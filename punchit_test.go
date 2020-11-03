package punchit

import (
	"strings"
	"testing"
)

func TestPunchCharValid(t *testing.T) {
	test := "A"
	exp := "100100000000"
	col, err := PunchChar(test)

	if err != nil {
		t.Errorf("Unexpected error:\n%v", err)
	}
	if len(col) != 12 {
		t.Errorf("Expected length 12, but got %v", len(col))
	}
	if col != exp {
		t.Errorf("Expected '%s', but got %v", exp, col)
	}
}

func TestPunchCharInvalid(t *testing.T) {
	test := "["
	_, err := PunchChar(test)

	if err == nil {
		t.Errorf("Expected error for invalid character '%s', but got nil", test)
	}
}

func TestPunchLine(t *testing.T) {
	test := "hello world"
	card, err := PunchLine(test)

	if err != nil {
		t.Errorf("Unexpected error:\n%v", err)
	}
	if strings.TrimSpace(card.src) != test {
		t.Errorf("Expected '%s', but got %v", test, card.src)
	}
	t.Logf(card.String())

	for i, col := range card.punches {
		if len(col) != 12 {
			t.Errorf("Col %d: Expected length 12, but got %d", i, len(col))
		}
	}
	for i := 0; i < 12; i++ {
		row := card.GetRow(i)
		if len(row) != 80 {
			t.Errorf("Row %d: Expected length 80, but got %d", i, len(row))
		}
	}
}

func TestPunchLines(t *testing.T) {
	test := []string{"first card", "second card", "third card"}
	deck, err := PunchLines(test)

	if err != nil {
		t.Errorf("Unexpected error:\n%v", err)
	}
	if len(deck) != 3 {
		t.Errorf("Expected deck of length 3, but got %v", len(deck))
	}
}
