package domain

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

type grammar struct {
	hash     hash.Hash
	root     Token
	channels []Channel
}

func createGrammar(
	hash hash.Hash,
	root Token,
) Grammar {
	return createGrammarInternally(hash, root, nil)
}

func createGrammarWithChannels(
	hash hash.Hash,
	root Token,
	channels []Channel,
) Grammar {
	return createGrammarInternally(hash, root, channels)
}

func createGrammarInternally(
	hash hash.Hash,
	root Token,
	channels []Channel,
) Grammar {
	out := grammar{
		hash:     hash,
		root:     root,
		channels: channels,
	}

	return &out
}

// Hash returns the hash
func (obj *grammar) Hash() hash.Hash {
	return obj.hash
}

// Root returns the root
func (obj *grammar) Root() Token {
	return obj.root
}

// HasChannels returns true if there is a channel, false otehrwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *grammar) Channels() []Channel {
	return obj.channels
}
