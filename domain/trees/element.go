package trees

import (
	"errors"
	"fmt"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type element struct {
	hash     hash.Hash
	contents []Content
	grammar  grammars.Element
}

func createElement(
	hash hash.Hash,
	contents []Content,
) Element {
	return createElementInternally(hash, contents, nil)
}

func createElementWithGrammar(
	hash hash.Hash,
	contents []Content,
	grammar grammars.Element,
) Element {
	return createElementInternally(hash, contents, grammar)
}

func createElementInternally(
	hash hash.Hash,
	contents []Content,
	grammar grammars.Element,
) Element {
	out := element{
		hash:     hash,
		grammar:  grammar,
		contents: contents,
	}

	return &out
}

// Fetch fetches a tree or value by name
func (obj *element) Fetch(hash hash.Hash, elementIndex uint) (Tree, Element, error) {
	if obj.HasGrammar() {
		if obj.grammar.Hash().Compare(hash) {
			return nil, obj, nil
		}
	}

	for _, oneContent := range obj.contents {
		if !oneContent.IsTree() {
			continue
		}

		tree, element, err := oneContent.Tree().Fetch(hash, elementIndex)
		if err != nil {
			continue
		}

		if tree != nil {
			return tree, nil, nil
		}

		if element != nil {
			return nil, element, nil
		}
	}

	str := fmt.Sprintf("there is no Tree or Element associated to the given hash: %s", hash.String())
	return nil, nil, errors.New(str)
}

// Bytes returns the element's bytes
func (obj *element) Bytes(includeChannels bool) []byte {
	output := []byte{}
	for _, oneContent := range obj.contents {
		if oneContent.IsValue() {
			value := oneContent.Value()
			if includeChannels && value.HasPrefix() {
				output = append(output, value.Prefix().Bytes(includeChannels)...)
			}

			output = append(output, value.Content()...)
			continue
		}

		output = append(output, oneContent.Tree().Bytes(includeChannels)...)
	}

	return output
}

// IsSuccessful returns true if successful, false otherwise
func (obj *element) IsSuccessful() bool {
	if !obj.HasGrammar() {
		return true
	}

	amount := obj.Amount()
	cardinality := obj.grammar.Cardinality()
	min := cardinality.Min()
	if amount < min {
		return false
	}

	if cardinality.HasMax() {
		pMax := cardinality.Max()
		if amount > *pMax {
			return false
		}
	}

	return true
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Contents returns the contents
func (obj *element) Contents() []Content {
	return obj.contents
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Element {
	return obj.grammar
}

// HasGrammar returns true if there is a grammar, false otherwise
func (obj *element) HasGrammar() bool {
	return obj.grammar != nil
}

// Amount returns the amount
func (obj *element) Amount() uint {
	return uint(len(obj.contents))
}
