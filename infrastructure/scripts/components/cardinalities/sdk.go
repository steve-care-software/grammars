package cardinalities

import grammars "github.com/steve-care-software/grammars/domain"

const byteLength = 256

// NewCardinality creates a new cardinality
func NewCardinality() Cardinality {
	cardinalityBuilder := grammars.NewCardinalityBuilder()
	return createCardinality(cardinalityBuilder)
}

// Cardinality represents the cardinality grammar
type Cardinality interface {
	Once() grammars.Cardinality
	Cardinality(min uint, pMax *uint) grammars.Cardinality
}
