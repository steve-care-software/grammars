package trees

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type valueBuilder struct {
	hashAdapter hash.Adapter
	content     []byte
	prefix      Trees
}

func createValueBuilder(
	hashAdapter hash.Adapter,
) ValueBuilder {
	out := valueBuilder{
		hashAdapter: hashAdapter,
		content:     nil,
		prefix:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder(
		app.hashAdapter,
	)
}

// WithContent adds a content to the builder
func (app *valueBuilder) WithContent(content []byte) ValueBuilder {
	app.content = content
	return app
}

// WithPrefix adds a prefix to the builder
func (app *valueBuilder) WithPrefix(prefix Trees) ValueBuilder {
	app.prefix = prefix
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.content != nil && len(app.content) <= 0 {
		app.content = nil
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Value instance")
	}

	data := [][]byte{
		app.content,
	}

	if app.prefix != nil {
		data = append(data, app.prefix.Bytes(false))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.prefix != nil {
		return createValueWithPrefix(*pHash, app.content, app.prefix), nil
	}

	return createValue(*pHash, app.content), nil
}
