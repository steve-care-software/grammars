package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	valid       []byte
	invalid     []byte
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		valid:       nil,
		invalid:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder(
		app.hashAdapter,
	)
}

// WithValid add valid bytes to the builder
func (app *suiteBuilder) WithValid(valid []byte) SuiteBuilder {
	app.valid = valid
	return app
}

// WithInvalid add invalid bytes to the builder
func (app *suiteBuilder) WithInvalid(invalid []byte) SuiteBuilder {
	app.invalid = invalid
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.valid != nil {
		pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
			[]byte("valid"),
			app.valid,
		})

		if err != nil {
			return nil, err
		}

		return createSuiteWithValid(*pHash, app.valid), nil
	}

	if app.invalid != nil {
		pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
			[]byte("invalid"),
			app.invalid,
		})

		if err != nil {
			return nil, err
		}

		return createSuiteWithInvalid(*pHash, app.invalid), nil
	}

	return nil, errors.New("the Suite is invalid")

}
