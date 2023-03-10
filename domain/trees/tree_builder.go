package trees

import (
	"errors"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type treeBuilder struct {
	hashAdapter hash.Adapter
	grammar     grammars.Token
	token       Token
	suffix      Trees
	remaining   []byte
}

func createTreeBuilder(
	hashAdapter hash.Adapter,
) TreeBuilder {
	out := treeBuilder{
		hashAdapter: hashAdapter,
		grammar:     nil,
		token:       nil,
		suffix:      nil,
		remaining:   nil,
	}

	return &out
}

// Create initializes the treeBuilder
func (app *treeBuilder) Create() TreeBuilder {
	return createTreeBuilder(
		app.hashAdapter,
	)
}

// WithGrammar adds a grammar to the treeBuilder
func (app *treeBuilder) WithGrammar(grammar grammars.Token) TreeBuilder {
	app.grammar = grammar
	return app
}

// WithToken adds a token to the treeBuilder
func (app *treeBuilder) WithToken(token Token) TreeBuilder {
	app.token = token
	return app
}

// WithSuffix adds a suffix to the builder
func (app *treeBuilder) WithSuffix(suffix Trees) TreeBuilder {
	app.suffix = suffix
	return app
}

// WithRemaining adds a remaining to the builder
func (app *treeBuilder) WithRemaining(remaining []byte) TreeBuilder {
	app.remaining = remaining
	return app
}

// Now builds a new Tree instance
func (app *treeBuilder) Now() (Tree, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Tree instance")
	}

	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Tree instance")
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = nil
	}

	data := [][]byte{
		app.grammar.Hash().Bytes(),
		app.token.Hash().Bytes(),
	}

	if app.suffix != nil {
		data = append(data, app.suffix.Hash().Bytes())
	}

	if app.remaining != nil {
		data = append(data, app.remaining)
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.remaining != nil && app.suffix != nil {
		return createTreeWithSuffixAndRemaining(*pHash, app.grammar, app.token, app.suffix, app.remaining), nil
	}

	if app.remaining != nil {
		return createTreeWithRemaining(*pHash, app.grammar, app.token, app.remaining), nil
	}

	if app.suffix != nil {
		return createTreeWithSuffix(*pHash, app.grammar, app.token, app.suffix), nil
	}

	return createTree(*pHash, app.grammar, app.token), nil
}
