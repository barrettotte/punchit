package main

import (
	"fmt"
	"strings"
)

// punch card dimensions (IBM 5081)
const cols = 80
const rows = 12

type punchcard struct {
	src   string           // source listing at top of punchcard
	holes [rows][cols]byte // store state of holes for each punchcard
}

func (p punchcard) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("     %s", strings.Repeat("_", cols+2)))
	sb.WriteString(fmt.Sprintf("\n     / %*s |", -cols, p.src))
	sb.WriteString(fmt.Sprintf("\n12  /  %s |", strings.Repeat(" ", cols)))
	sb.WriteString(fmt.Sprintf("\n11 |  %s  |", strings.Repeat(" ", cols)))

	for i := 0; i < 10; i++ {
		sb.WriteString(fmt.Sprintf("\n %d |  %s  |", i, strings.Repeat(" ", cols)))
	}
	sb.WriteString(fmt.Sprintf("\n   |%s|", strings.Repeat("_", cols+4)))

	return sb.String()
}
