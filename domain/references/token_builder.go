package references

import (
	"errors"

	grammars "github.com/steve-care-software/grammars/domain"
)

type tokenBuilder struct {
	name      string
	reference grammars.Token
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:      "",
		reference: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithReference adds a reference to the builder
func (app *tokenBuilder) WithReference(reference grammars.Token) TokenBuilder {
	app.reference = reference
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Token instance")
	}

	return createToken(app.name, app.reference), nil
}
