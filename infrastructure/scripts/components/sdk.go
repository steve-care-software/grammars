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

// NewComponent creates a new component
func NewComponent() Component {
	cardinality := cardinalities.NewCardinality()
	channel := channels.NewChannel()
	suite := suites.NewSuite()
	element := elements.NewElement()
	line := lines.NewLine()
	everything := everythings.NewEverything()
	token := tokens.NewToken()
	return createComponent(
		cardinality,
		channel,
		suite,
		element,
		line,
		everything,
		token,
	)
}

// Component represents the component
type Component interface {
	Cardinality() cardinalities.Cardinality
	Channel() channels.Channel
	Suite() suites.Suite
	Element() elements.Element
	Line() lines.Line
	Everything() everythings.Everything
	Token() tokens.Token
}
