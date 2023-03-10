package references

import grammars "github.com/steve-care-software/grammars/domain"

type reference struct {
	root     grammars.Grammar
	tokens   Tokens
	grammars Grammars
}

func createReference(
	root grammars.Grammar,
	tokens Tokens,
) Reference {
	return createReferenceInternally(root, tokens, nil)
}

func createReferenceWithGrammars(
	root grammars.Grammar,
	tokens Tokens,
	grammars Grammars,
) Reference {
	return createReferenceInternally(root, tokens, grammars)
}

func createReferenceInternally(
	root grammars.Grammar,
	tokens Tokens,
	grammars Grammars,
) Reference {
	out := reference{
		root:     root,
		tokens:   tokens,
		grammars: grammars,
	}

	return &out
}

// Root returns the root grammar
func (obj *reference) Root() grammars.Grammar {
	return obj.root
}

// Tokens returns the tokens
func (obj *reference) Tokens() Tokens {
	return obj.tokens
}

// HasGrammars returns true if there is grammars, false otherwise
func (obj *reference) HasGrammars() bool {
	return obj.grammars != nil
}

// Grammars returns the grammars, if any
func (obj *reference) Grammars() Grammars {
	return obj.grammars
}
