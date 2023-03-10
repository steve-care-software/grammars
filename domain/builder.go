package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	root        Token
	channels    []Channel
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		root:        nil,
		channels:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root Token) Builder {
	app.root = root
	return app
}

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels []Channel) Builder {
	app.channels = channels
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a Grammar instance")
	}

	if app.channels != nil && len(app.channels) <= 0 {
		app.channels = nil
	}

	data := [][]byte{
		app.root.Hash().Bytes(),
	}

	if app.channels != nil {
		for _, oneChannel := range app.channels {
			data = append(data, oneChannel.Hash().Bytes())
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.channels != nil {
		return createGrammarWithChannels(*pHash, app.root, app.channels), nil
	}

	return createGrammar(*pHash, app.root), nil
}
