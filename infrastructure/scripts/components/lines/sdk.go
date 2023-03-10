package lines

import grammars "github.com/steve-care-software/grammars/domain"

// NewLine creates a new line
func NewLine() Line {
	lineBuilder := grammars.NewLineBuilder()
	return createLine(
		lineBuilder,
	)
}

// Line represents a line component
type Line interface {
	FromElements(elements []grammars.Element) grammars.Line
}
