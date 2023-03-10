package trees

import "github.com/steve-care-software/libs/cryptography/hash"

type token struct {
	hash       hash.Hash
	lines      []Line
	successful Line
}

func createToken(
	hash hash.Hash,
	lines []Line,
) Token {
	return createTokenInternally(hash, lines, nil)
}

func createTokenWithSuccessful(
	hash hash.Hash,
	lines []Line,
	successful Line,
) Token {
	return createTokenInternally(hash, lines, successful)
}

func createTokenInternally(
	hash hash.Hash,
	lines []Line,
	successful Line,
) Token {
	out := token{
		hash:       hash,
		lines:      lines,
		successful: successful,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Lines returns the lines
func (obj *token) Lines() []Line {
	return obj.lines
}

// HasSuccessful returns true if there is a successful line, false otherwise
func (obj *token) HasSuccessful() bool {
	return obj.successful != nil
}

// Successful returns the successful line, if any
func (obj *token) Successful() Line {
	return obj.successful
}
