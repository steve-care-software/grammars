package trees

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	lines       []Line
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		lines:       nil,
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

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Token instance")
	}

	data := [][]byte{}
	for _, oneLine := range app.lines {
		data = append(data, oneLine.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	var successful Line
	for _, oneLine := range app.lines {
		if oneLine.IsSuccessful() {
			successful = oneLine
			break
		}
	}

	if successful != nil {
		return createTokenWithSuccessful(*pHash, app.lines, successful), nil
	}

	return createToken(*pHash, app.lines), nil
}
