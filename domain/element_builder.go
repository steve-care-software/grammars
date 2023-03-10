package domain

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	cardinality Cardinality
	value       []byte
	grammar     Grammar
	instance    Instance
	recursive   string
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		cardinality: nil,
		value:       nil,
		grammar:     nil,
		instance:    nil,
		recursive:   "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithCardinality adds a cardinality to the builder
func (app *elementBuilder) WithCardinality(cardinality Cardinality) ElementBuilder {
	app.cardinality = cardinality
	return app
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value []byte) ElementBuilder {
	app.value = value
	return app
}

// WithGrammar adds an grammar grammar to the builder
func (app *elementBuilder) WithGrammar(grammar Grammar) ElementBuilder {
	app.grammar = grammar
	return app
}

// WithInstance adds an instance to the builder
func (app *elementBuilder) WithInstance(instance Instance) ElementBuilder {
	app.instance = instance
	return app
}

// WithRecursive adds a recursive to the builder
func (app *elementBuilder) WithRecursive(recursive string) ElementBuilder {
	app.recursive = recursive
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build an Element instance")
	}

	if app.value != nil && len(app.value) <= 0 {
		app.value = nil
	}

	contentData := [][]byte{}
	if app.value != nil {
		contentData = append(contentData, app.value)
	}

	if app.grammar != nil {
		contentData = append(contentData, app.grammar.Hash())
	}

	if app.instance != nil {
		contentData = append(contentData, app.instance.Hash())
	}

	if app.recursive != "" {
		contentData = append(contentData, []byte(app.recursive))
	}

	if len(contentData) <= 0 {

	}

	pContentHash, err := app.hashAdapter.FromMultiBytes(contentData)
	if err != nil {
		return nil, err
	}

	data := [][]byte{
		pContentHash.Bytes(),
		[]byte(fmt.Sprintf("%d", app.cardinality.Min())),
	}

	if app.cardinality.HasMax() {
		pMax := app.cardinality.Max()
		data = append(data, []byte(fmt.Sprintf("%d", *pMax)))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.value != nil {
		content := createElementContentWithValue(*pContentHash, app.value)
		return createElement(*pHash, content, app.cardinality), nil
	}

	if app.grammar != nil {
		content := createElementContentWithGrammar(*pContentHash, app.grammar)
		return createElement(*pHash, content, app.cardinality), nil
	}

	if app.instance != nil {
		content := createElementContentWithInstance(*pContentHash, app.instance)
		return createElement(*pHash, content, app.cardinality), nil
	}

	if app.recursive != "" {
		content := createElementContentWithRecursive(*pContentHash, app.recursive)
		return createElement(*pHash, content, app.cardinality), nil
	}

	return nil, errors.New("the Element is invalid")
}
