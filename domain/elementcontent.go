package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type elementContent struct {
	hash      hash.Hash
	value     []byte
	grammar      Grammar
	instance  Instance
	recursive string
}

func createElementContentWithValue(
	hash hash.Hash,
	value []byte,
) ElementContent {
	return createElementContentInternally(hash, value, nil, nil, "")
}

func createElementContentWithGrammar(
	hash hash.Hash,
	grammar Grammar,
) ElementContent {
	return createElementContentInternally(hash, nil, grammar, nil, "")
}

func createElementContentWithInstance(
	hash hash.Hash,
	instance Instance,
) ElementContent {
	return createElementContentInternally(hash, nil, nil, instance, "")
}

func createElementContentWithRecursive(
	hash hash.Hash,
	recursive string,
) ElementContent {
	return createElementContentInternally(hash, nil, nil, nil, recursive)
}

func createElementContentInternally(
	hash hash.Hash,
	value []byte,
	grammar Grammar,
	instance Instance,
	recursive string,
) ElementContent {
	out := elementContent{
		hash:      hash,
		value:     value,
		grammar:      grammar,
		instance:  instance,
		recursive: recursive,
	}

	return &out
}

// Hash return the hash
func (obj *elementContent) Hash() hash.Hash {
	return obj.hash
}

// IsValue returns true if there is a value, false otherwise
func (obj *elementContent) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *elementContent) Value() []byte {
	return obj.value
}

// IsGrammar returns true if there is an grammar grammar, false otherwise
func (obj *elementContent) IsGrammar() bool {
	return obj.grammar != nil
}

// Grammar returns the grammar grammar, if any
func (obj *elementContent) Grammar() Grammar {
	return obj.grammar
}

// IsInstance returns true if there is an instance, false otherwise
func (obj *elementContent) IsInstance() bool {
	return obj.instance != nil
}

// Instance returns the instance, if any
func (obj *elementContent) Instance() Instance {
	return obj.instance
}

// IsRecursive returns true if there is a recursive token, false otherwise
func (obj *elementContent) IsRecursive() bool {
	return obj.recursive != ""
}

// Recursive returns the recursive, if any
func (obj *elementContent) Recursive() string {
	return obj.recursive
}
