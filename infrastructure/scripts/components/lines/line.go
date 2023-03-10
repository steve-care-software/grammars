package lines

import grammars "github.com/steve-care-software/grammars/domain"

type line struct {
	lineBuilder grammars.LineBuilder
}

func createLine(
	lineBuilder grammars.LineBuilder,
) Line {
	out := line{
		lineBuilder: lineBuilder,
	}

	return &out
}

// FromElements returns a line from elements
func (app *line) FromElements(elements []grammars.Element) grammars.Line {
	return app.lineFromElements(elements)
}

func (app *line) lineFromElements(elements []grammars.Element) grammars.Line {
	line, err := app.lineBuilder.Create().
		WithElements(elements).
		Now()

	if err != nil {
		panic(err)
	}

	return line
}
