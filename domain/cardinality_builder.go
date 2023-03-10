package domain

import (
	"errors"
)

type cardinalityBuilder struct {
	pMin *uint
	pMax *uint
}

func createCardinalityBuilder() CardinalityBuilder {
	out := cardinalityBuilder{
		pMin: nil,
		pMax: nil,
	}

	return &out
}

// Create initializes the builder
func (app *cardinalityBuilder) Create() CardinalityBuilder {
	return createCardinalityBuilder()
}

// WithMin adds a minimum to the builder
func (app *cardinalityBuilder) WithMin(min uint) CardinalityBuilder {
	app.pMin = &min
	return app
}

// WithMax adds a maximum to the builder
func (app *cardinalityBuilder) WithMax(max uint) CardinalityBuilder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *cardinalityBuilder) Now() (Cardinality, error) {
	if app.pMin == nil {
		return nil, errors.New("the minimum is mandatory in order to build a Cardinality instance")
	}

	if app.pMax != nil {
		return createCardinalityWithMax(*app.pMin, app.pMax), nil
	}

	return createCardinality(*app.pMin), nil
}
