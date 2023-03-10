package references

import "errors"

type grammarsBuilder struct {
	list []Grammar
}

func createGrammarsBuilder() GrammarsBuilder {
	out := grammarsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *grammarsBuilder) Create() GrammarsBuilder {
	return createGrammarsBuilder()
}

// WithList adds a list to the builder
func (app *grammarsBuilder) WithList(list []Grammar) GrammarsBuilder {
	app.list = list
	return app
}

// Now builds a new Grammars instance
func (app *grammarsBuilder) Now() (Grammars, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Grammar in order to build a Grammars instance")
	}

	mp := map[string]Grammar{}
	for _, oneGrammar := range app.list {
		keyname := oneGrammar.Reference().Hash().String()
		mp[keyname] = oneGrammar
	}

	return createGrammars(app.list, mp), nil
}
