package cardinalities

import grammars "github.com/steve-care-software/grammars/domain"

type cardinality struct {
	cardinalityBuilder grammars.CardinalityBuilder
}

func createCardinality(
	cardinalityBuilder grammars.CardinalityBuilder,
) Cardinality {
	out := cardinality{
		cardinalityBuilder: cardinalityBuilder,
	}

	return &out
}

// Once returns acardinality with 1 element
func (app *cardinality) Once() grammars.Cardinality {
	max := uint(1)
	return app.Cardinality(1, &max)
}

// Cardinality returns a cardinality instance
func (app *cardinality) Cardinality(min uint, pMax *uint) grammars.Cardinality {
	builder := app.cardinalityBuilder.Create().WithMin(min)
	if pMax != nil {
		builder.WithMax(*pMax)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
