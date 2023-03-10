package tokens

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	script_elements "github.com/steve-care-software/grammars/infrastructure/scripts/components/elements"
	script_lines "github.com/steve-care-software/grammars/infrastructure/scripts/components/lines"
	script_suites "github.com/steve-care-software/grammars/infrastructure/scripts/components/suites"
)

const byteLength = 256

// NewToken creates a new token instance
func NewToken() Token {
	suite := script_suites.NewSuite()
	element := script_elements.NewElement()
	line := script_lines.NewLine()
	tokenBuilder := grammars.NewTokenBuilder()
	refTokenBuilder := references.NewTokenBuilder()
	return createToken(
		suite,
		element,
		line,
		tokenBuilder,
		refTokenBuilder,
	)
}

// Token represents the token component
type Token interface {
	AllCharacters(tokenName string, values string) references.Token
	AnyCharacter(tokenName string, values string) references.Token
	AnyElement(tokenName string, elementsList []grammars.Element, suites []grammars.Suite) references.Token
	FromLines(name string, lines []grammars.Line, suites []grammars.Suite) references.Token
}
