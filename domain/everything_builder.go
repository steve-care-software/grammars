package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type everythingBuilder struct {
	hashAdapter hash.Adapter
	exception   Token
	escape      Token
}

func createEverythingBuilder(
	hashAdapter hash.Adapter,
) EverythingBuilder {
	out := everythingBuilder{
		hashAdapter: hashAdapter,
		exception:   nil,
		escape:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *everythingBuilder) Create() EverythingBuilder {
	return createEverythingBuilder(
		app.hashAdapter,
	)
}

// WithException adds an exception to the builder
func (app *everythingBuilder) WithException(exception Token) EverythingBuilder {
	app.exception = exception
	return app
}

// WithEscape adds an escape to the builder
func (app *everythingBuilder) WithEscape(escape Token) EverythingBuilder {
	app.escape = escape
	return app
}

// Now builds a new Everything instance
func (app *everythingBuilder) Now() (Everything, error) {
	if app.exception == nil {
		return nil, errors.New("the exception is mandatory in order to build an Everything instance")
	}

	data := [][]byte{
		app.exception.Hash().Bytes(),
	}

	if app.escape != nil {
		data = append(data, app.escape.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.escape != nil {
		return createEverythingWithEscape(*pHash, app.exception, app.escape), nil
	}

	return createEverything(*pHash, app.exception), nil
}
