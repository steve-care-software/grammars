package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type channelBuilder struct {
	hashAdapter hash.Adapter
	token       Token
	condition   ChannelCondition
}

func createChannelBuilder(
	hashAdapter hash.Adapter,
) ChannelBuilder {
	out := channelBuilder{
		hashAdapter: hashAdapter,
		token:       nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelBuilder) Create() ChannelBuilder {
	return createChannelBuilder(
		app.hashAdapter,
	)
}

// WithToken adds a token to the builder
func (app *channelBuilder) WithToken(token Token) ChannelBuilder {
	app.token = token
	return app
}

// WithCondition adds a condition to the builder
func (app *channelBuilder) WithCondition(condition ChannelCondition) ChannelBuilder {
	app.condition = condition
	return app
}

// Now builds a new Channel instance
func (app *channelBuilder) Now() (Channel, error) {
	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Channel instance")
	}

	data := [][]byte{
		app.token.Hash().Bytes(),
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.condition != nil {
		return createChannelWithCondition(*pHash, app.token, app.condition), nil
	}

	return createChannel(*pHash, app.token), nil
}
