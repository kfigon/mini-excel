package main

import "fmt"

type tokenType int

const (
	number tokenType = iota
	plus
	minus
	multiply
	equal
	openParent
	closeParent
	coord
)

func (t tokenType) String() string {
	return []string{"number",
		"plus",
		"minus",
		"multiply",
		"equal",
		"openParent",
		"closeParent",
		"coord"}[t]
}

type token struct {
	tokType tokenType
	val     string
}

func (t token) String() string {
	return fmt.Sprintf("<%v; %v>", t.tokType, t.val)
}

func parseExpression(ex expressionCell) []token {
	return nil
}