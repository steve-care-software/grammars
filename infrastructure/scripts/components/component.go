package components

import (
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/cardinalities"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/channels"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/elements"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/everythings"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/lines"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/suites"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components/tokens"
)

type component struct {
	cardinality cardinalities.Cardinality
	channel     channels.Channel
	suite       suites.Suite
	element     elements.Element
	line        lines.Line
	everything  everythings.Everything
	token       tokens.Token
}

func createComponent(
	cardinality cardinalities.Cardinality,
	channel channels.Channel,
	suite suites.Suite,
	element elements.Element,
	line lines.Line,
	everything everythings.Everything,
	token tokens.Token,
) Component {
	out := component{
		cardinality: cardinality,
		channel:     channel,
		suite:       suite,
		element:     element,
		line:        line,
		everything:  everything,
		token:       token,
	}

	return &out
}

// Cardinality returns the cardinality
func (app *component) Cardinality() cardinalities.Cardinality {
	return app.cardinality
}

// Channel returns the channel
func (app *component) Channel() channels.Channel {
	return app.channel
}

// Suite returns the suite
func (app *component) Suite() suites.Suite {
	return app.suite
}

// Element returns the element
func (app *component) Element() elements.Element {
	return app.element
}

// Line returns the line
func (app *component) Line() lines.Line {
	return app.line
}

// Everything returns the everything
func (app *component) Everything() everythings.Everything {
	return app.everything
}

// Token returns the token
func (app *component) Token() tokens.Token {
	return app.token
}
