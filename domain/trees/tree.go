package trees

import (
	"errors"
	"fmt"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type tree struct {
	hash      hash.Hash
	grammar   grammars.Token
	token     Token
	suffix    Trees
	remaining []byte
}

func createTree(
	hash hash.Hash,
	grammar grammars.Token,
	token Token,
) Tree {
	return createTreeInternally(hash, grammar, token, nil, nil)
}

func createTreeWithRemaining(
	hash hash.Hash,
	grammar grammars.Token,
	token Token,
	remaining []byte,
) Tree {
	return createTreeInternally(hash, grammar, token, nil, remaining)
}

func createTreeWithSuffix(
	hash hash.Hash,
	grammar grammars.Token,
	token Token,
	suffix Trees,
) Tree {
	return createTreeInternally(hash, grammar, token, suffix, nil)
}

func createTreeWithSuffixAndRemaining(
	hash hash.Hash,
	grammar grammars.Token,
	token Token,
	suffix Trees,
	remaining []byte,
) Tree {
	return createTreeInternally(hash, grammar, token, suffix, remaining)
}

func createTreeInternally(
	hash hash.Hash,
	grammar grammars.Token,
	token Token,
	suffix Trees,
	remaining []byte,
) Tree {
	out := tree{
		hash:      hash,
		grammar:   grammar,
		token:     token,
		suffix:    suffix,
		remaining: remaining,
	}

	return &out
}

// Fetch fetches a tree or value by name
func (obj *tree) Fetch(hash hash.Hash, elementIndex uint) (Tree, Element, error) {
	if obj.grammar.Hash().Compare(hash) {
		return obj, nil, nil
	}

	str := fmt.Sprintf("there is no Tree or Element associated to the given hash: %s,at element's index: %d", hash, elementIndex)
	if !obj.Token().HasSuccessful() {
		return nil, nil, errors.New(str)
	}

	cpt := uint(0)
	elementsList := obj.Token().Successful().Elements()
	for _, oneElement := range elementsList {
		tree, element, err := oneElement.Fetch(hash, elementIndex)
		if err != nil {
			continue
		}

		isReady := cpt >= elementIndex
		if tree != nil && isReady {
			return tree, nil, nil
		}

		if element != nil && isReady {
			return nil, element, nil
		}

		if tree != nil || element != nil {
			cpt++
		}
	}

	return nil, nil, errors.New(str)
}

// Bytes returns the tree's bytes
func (obj *tree) Bytes(includeChannels bool) []byte {
	output := []byte{}
	if !obj.token.HasSuccessful() {
		return output
	}

	elements := obj.token.Successful().Elements()
	for _, oneElement := range elements {
		output = append(output, oneElement.Bytes(includeChannels)...)
	}

	if includeChannels && obj.HasSuffix() {
		output = append(output, obj.Suffix().Bytes(includeChannels)...)
	}

	return output
}

// Hash returns the hash, if any
func (obj *tree) Hash() hash.Hash {
	return obj.hash
}

// Grammar returns the grammar
func (obj *tree) Grammar() grammars.Token {
	return obj.grammar
}

// Token returns the token
func (obj *tree) Token() Token {
	return obj.token
}

// HasSuffix returns true if there is suffix, false otherwise
func (obj *tree) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the token
func (obj *tree) Suffix() Trees {
	return obj.suffix
}

// HasRemaining returns true if there is remaining, false otherwise
func (obj *tree) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns remaining, if any
func (obj *tree) Remaining() []byte {
	return obj.remaining
}
