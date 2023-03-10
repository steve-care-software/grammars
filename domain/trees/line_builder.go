package trees

import (
	"errors"
	"fmt"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type lineBuilder struct {
	hashAdapter hash.Adapter
	pIndex      *uint
	grammar     grammars.Line
	isReverse   bool
	elements    []Element
}

func createLineBuilder(
	hashAdapter hash.Adapter,
) LineBuilder {
	out := lineBuilder{
		hashAdapter: hashAdapter,
		pIndex:      nil,
		grammar:     nil,
		isReverse:   false,
		elements:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *lineBuilder) WithIndex(index uint) LineBuilder {
	app.pIndex = &index
	return app
}

// WithGrammar adds a grammar to the builder
func (app *lineBuilder) WithGrammar(grammar grammars.Line) LineBuilder {
	app.grammar = grammar
	return app
}

// WithElements add elements to the builder
func (app *lineBuilder) WithElements(elements []Element) LineBuilder {
	app.elements = elements
	return app
}

// IsReverse flags the builder as reverse
func (app *lineBuilder) IsReverse() LineBuilder {
	app.isReverse = true
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Line instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Line instance")
	}

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", *app.pIndex)),
		app.grammar.Hash().Bytes(),
	}

	if app.elements != nil {
		for _, oneElement := range app.elements {
			data = append(data, oneElement.Hash().Bytes())
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.elements != nil {
		mp := map[string]Element{}
		for _, oneElement := range app.elements {
			if !oneElement.HasGrammar() {
				continue
			}

			keyname := oneElement.Grammar().Hash().String()
			mp[keyname] = oneElement
		}

		return createLineWithElements(*pHash, *app.pIndex, app.grammar, app.isReverse, app.elements, mp), nil
	}

	return createLine(*pHash, *app.pIndex, app.grammar, app.isReverse), nil
}
