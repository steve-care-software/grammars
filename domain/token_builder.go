package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	lines       []Line
	suites      []Suite
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		lines:       nil,
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder(
		app.hashAdapter,
	)
}

// WithLines add lines to the builder
func (app *tokenBuilder) WithLines(lines []Line) TokenBuilder {
	app.lines = lines
	return app
}

// WithSuites add suites to the builder
func (app *tokenBuilder) WithSuites(suites []Suite) TokenBuilder {
	app.suites = suites
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Token instance")
	}

	if app.suites != nil && len(app.suites) <= 0 {
		app.suites = nil
	}

	data := [][]byte{}
	for _, oneLine := range app.lines {
		data = append(data, oneLine.Hash().Bytes())
	}

	if app.suites != nil {
		for _, oneSuite := range app.suites {
			data = append(data, oneSuite.Hash().Bytes())
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createTokenWithSuites(*pHash, app.lines, app.suites), nil
	}

	return createToken(*pHash, app.lines), nil
}
