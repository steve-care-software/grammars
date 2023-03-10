package channels

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	script_cardinalities "github.com/steve-care-software/grammars/infrastructure/scripts/components/cardinalities"
	script_elements "github.com/steve-care-software/grammars/infrastructure/scripts/components/elements"
	script_everythings "github.com/steve-care-software/grammars/infrastructure/scripts/components/everythings"
	script_lines "github.com/steve-care-software/grammars/infrastructure/scripts/components/lines"
	script_suites "github.com/steve-care-software/grammars/infrastructure/scripts/components/suites"
	script_tokens "github.com/steve-care-software/grammars/infrastructure/scripts/components/tokens"
)

// NewChannel creates a new channel component
func NewChannel() Channel {
	token := script_tokens.NewToken()
	element := script_elements.NewElement()
	line := script_lines.NewLine()
	suite := script_suites.NewSuite()
	everything := script_everythings.NewEverything()
	cardinality := script_cardinalities.NewCardinality()
	channelBuilder := grammars.NewChannelBuilder()
	return createChannel(
		token,
		element,
		line,
		suite,
		everything,
		cardinality,
		channelBuilder,
	)
}

// Channel represents the channel grammar
type Channel interface {
	Channels() ([]grammars.Channel, []references.Token)
}
