package elements

import (
	grammars "github.com/steve-care-software/grammars/domain"
	element_cardinalities "github.com/steve-care-software/grammars/infrastructure/scripts/components/cardinalities"
)

type element struct {
	cardinality     element_cardinalities.Cardinality
	instanceBuilder grammars.InstanceBuilder
	elementBuilder  grammars.ElementBuilder
}

func createElement(
	cardinality element_cardinalities.Cardinality,
	instanceBuilder grammars.InstanceBuilder,
	elementBuilder grammars.ElementBuilder,
) Element {
	out := element{
		cardinality:     cardinality,
		instanceBuilder: instanceBuilder,
		elementBuilder:  elementBuilder,
	}

	return &out
}

// FromEverything returns the element from everything
func (app *element) FromEverything(everything grammars.Everything) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithEverything(everything).
		Now()

	if err != nil {
		panic(err)
	}

	cardinality := app.cardinality.Once()
	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

// FromToken retruns an element from token
func (app *element) FromToken(token grammars.Token, cardinality grammars.Cardinality) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithToken(token).
		Now()

	if err != nil {
		panic(err)
	}

	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

// FromValue creates an element from value
func (app *element) FromValue(value []byte) grammars.Element {
	ins, err := app.elementBuilder.Create().
		WithValue(value).
		WithCardinality(app.cardinality.Once()).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
