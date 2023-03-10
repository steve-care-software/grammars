package trees

import (
	"errors"
	"fmt"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type line struct {
	hash      hash.Hash
	index     uint
	grammar   grammars.Line
	isReverse bool
	elements  []Element
	mp        map[string]Element
}

func createLine(
	hash hash.Hash,
	index uint,
	grammar grammars.Line,
	isReverse bool,
) Line {
	return createLineInternally(hash, index, grammar, isReverse, nil, map[string]Element{})
}

func createLineWithElements(
	hash hash.Hash,
	index uint,
	grammar grammars.Line,
	isReverse bool,
	elements []Element,
	mp map[string]Element,
) Line {
	return createLineInternally(hash, index, grammar, isReverse, elements, mp)
}

func createLineInternally(
	hash hash.Hash,
	index uint,
	grammar grammars.Line,
	isReverse bool,
	elements []Element,
	mp map[string]Element,
) Line {
	out := line{
		hash:      hash,
		index:     index,
		grammar:   grammar,
		isReverse: isReverse,
		elements:  elements,
		mp:        mp,
	}

	return &out
}

// Hash returns the hash
func (obj *line) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *line) Index() uint {
	return obj.index
}

// IsReverse returns true if reverse, false otherwise
func (obj *line) IsReverse() bool {
	return obj.isReverse
}

// Grammar returns the grammar
func (obj *line) Grammar() grammars.Line {
	return obj.grammar
}

// HasElements returns true if there is elements, false otherwise
func (obj *line) HasElements() bool {
	return obj.elements != nil
}

// Elements returns the elements
func (obj *line) Elements() []Element {
	return obj.elements
}

// Fetch fetches an element by name
func (obj *line) Fetch(hash hash.Hash) (Element, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the element (hash: %s) does not exists", keyname)
	return nil, errors.New(str)
}

// IsSuccessful returns true if successful, false otherwise
func (obj *line) IsSuccessful() bool {
	if !obj.HasElements() {
		return false
	}

	requested := obj.grammar.Elements()
	for _, oneElement := range obj.elements {
		if !oneElement.IsSuccessful() {
			return false
		}
	}

	if obj.IsReverse() {
		return true
	}

	for _, oneElement := range requested {
		requestedMin := oneElement.Cardinality().Min()
		if requestedMin <= 0 {
			continue
		}

		requestedHash := oneElement.Hash()
		element, err := obj.Fetch(requestedHash)
		if err != nil {
			return false
		}

		amount := element.Amount()
		if requestedMin > amount {
			return false
		}
	}

	return true
}
