package elements

import (
	grammars "github.com/steve-care-software/grammars/domain"
	element_cardinalities "github.com/steve-care-software/grammars/infrastructure/scripts/components/cardinalities"
)

// NewElement creates a new element instance
func NewElement() Element {
	cardinality := element_cardinalities.NewCardinality()
	instanceBuilder := grammars.NewInstanceBuilder()
	elementBuilder := grammars.NewElementBuilder()
	return createElement(
		cardinality,
		instanceBuilder,
		elementBuilder,
	)
}

// Element represents the element
type Element interface {
	FromEverything(everything grammars.Everything) grammars.Element
	FromToken(token grammars.Token, cardinality grammars.Cardinality) grammars.Element
	FromValue(value []byte) grammars.Element
}
