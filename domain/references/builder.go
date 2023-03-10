package references

import (
	"errors"

	grammars "github.com/steve-care-software/grammars/domain"
)

type builder struct {
	root     grammars.Grammar
	tokens   Tokens
	grammars Grammars
}

func createBuilder() Builder {
	out := builder{
		root:     nil,
		tokens:   nil,
		grammars: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root grammars.Grammar) Builder {
	app.root = root
	return app
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithGrammars add grammars to the builder
func (app *builder) WithGrammars(grammars Grammars) Builder {
	app.grammars = grammars
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	if app.root == nil {
		return nil, errors.New("the root grammar is mandatory in order to build a Reference instance")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens are mandatory in order to build a Reference instance")
	}

	if app.grammars != nil {
		return createReferenceWithGrammars(app.root, app.tokens, app.grammars), nil
	}

	return createReference(app.root, app.tokens), nil
}
