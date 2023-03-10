package everythings

import grammars "github.com/steve-care-software/grammars/domain"

// NewEverything creates a new everything component
func NewEverything() Everything {
	everythingBuilder := grammars.NewEverythingBuilder()
	return createEverything(everythingBuilder)
}

// Everything represents an everything component
type Everything interface {
	WithoutEscape(exception grammars.Token) grammars.Everything
	Everything(exception grammars.Token, escape grammars.Token) grammars.Everything
}
