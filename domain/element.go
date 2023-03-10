package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type element struct {
	hash        hash.Hash
	content     ElementContent
	cardinality Cardinality
}

func createElement(
	hash hash.Hash,
	content ElementContent,
	cardinality Cardinality,
) Element {
	out := element{
		hash:        hash,
		content:     content,
		cardinality: cardinality,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *element) Content() ElementContent {
	return obj.content
}

// Cardinality returns the cardinality
func (obj *element) Cardinality() Cardinality {
	return obj.cardinality
}
