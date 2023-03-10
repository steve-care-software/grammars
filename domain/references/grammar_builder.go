package references

import (
	"errors"

	grammars "github.com/steve-care-software/grammars/domain"
)

type grammarBuilder struct {
	name      string
	reference grammars.Grammar
}

func createGrammarBuilder() GrammarBuilder {
	out := grammarBuilder{
		name:      "",
		reference: nil,
	}

	return &out
}

// Create initializes the builder
func (app *grammarBuilder) Create() GrammarBuilder {
	return createGrammarBuilder()
}

// WithName adds a name to the builder
func (app *grammarBuilder) WithName(name string) GrammarBuilder {
	app.name = name
	return app
}

// WithReference adds a reference to the builder
func (app *grammarBuilder) WithReference(reference grammars.Grammar) GrammarBuilder {
	app.reference = reference
	return app
}

// Now builds a new Grammar instance
func (app *grammarBuilder) Now() (Grammar, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Grammar instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Grammar instance")
	}

	return createGrammar(app.name, app.reference), nil
}
