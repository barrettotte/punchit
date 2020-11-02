package main

import (
	"fmt"
	"strings"
)

// punch card dimensions (IBM 5081)
const cols = 80
const rows = 12

type punchcard struct {
	src     string       // source listing at top
	punches [cols]string // holes for each column
}

func (p *punchcard) String() string {
	var sb strings.Builder
	r := strings.NewReplacer("0", " ", "1", "▮")

	// top
	sb.WriteString(fmt.Sprintf("      %s", strings.Repeat("_", cols+2)))
	sb.WriteString(fmt.Sprintf("\n     / %*s |", -cols, p.src))

	// punches
	sb.WriteString(fmt.Sprintf("\n12  /  %s |", r.Replace(p.getRow(0))))
	sb.WriteString(fmt.Sprintf("\n11 |   %s |", r.Replace(p.getRow(1))))
	for i := 2; i < rows; i++ {
		sb.WriteString(fmt.Sprintf("\n %d |   %s |", i-2, r.Replace(p.getRow(i))))
	}

	// bottom
	sb.WriteString(fmt.Sprintf("\n   |%s|", strings.Repeat("_", cols+4)))

	return sb.String()
}

func (p *punchcard) getRow(i int) string {
	var row string

	for j := 0; j < cols; j++ {
		row += string(p.punches[j][i])
	}
	return row
}

func (p *punchcard) getCol(i int) string {
	return p.punches[i]
}
