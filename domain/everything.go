package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type everything struct {
	hash      hash.Hash
	exception Token
	escape    Token
}

func createEverything(
	hash hash.Hash,
	exception Token,
) Everything {
	return createEverythingInternally(hash, exception, nil)
}

func createEverythingWithEscape(
	hash hash.Hash,
	exception Token,
	escape Token,
) Everything {
	return createEverythingInternally(hash, exception, escape)
}

func createEverythingInternally(
	hash hash.Hash,
	exception Token,
	escape Token,
) Everything {
	out := everything{
		hash:      hash,
		exception: exception,
		escape:    escape,
	}

	return &out
}

// Hash returns the hash
func (obj *everything) Hash() hash.Hash {
	return obj.hash
}

// Exception returns the exception
func (obj *everything) Exception() Token {
	return obj.exception
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *everything) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *everything) Escape() Token {
	return obj.escape
}
