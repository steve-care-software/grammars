package trees

import (
	"errors"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	grammar     grammars.Element
	contents    []Content
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		grammar:     nil,
		contents:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithGrammar adds a grammar to the builder
func (app *elementBuilder) WithGrammar(grammar grammars.Element) ElementBuilder {
	app.grammar = grammar
	return app
}

// WithContents adds a contents to the builder
func (app *elementBuilder) WithContents(contents []Content) ElementBuilder {
	app.contents = contents
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.contents != nil && len(app.contents) <= 0 {
		app.contents = nil
	}

	if app.contents == nil {
		return nil, errors.New("the contents is mandatory in order to build an Element instance")
	}

	data := [][]byte{}
	for _, oneContent := range app.contents {
		data = append(data, oneContent.Hash().Bytes())
	}

	if app.grammar != nil {
		data = append(data, app.grammar.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.grammar != nil {
		return createElementWithGrammar(*pHash, app.contents, app.grammar), nil
	}

	return createElement(*pHash, app.contents), nil
}
