package references

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	return createTokensBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewGrammarsBuilder creates a new grammars builder
func NewGrammarsBuilder() GrammarsBuilder {
	return createGrammarsBuilder()
}

// NewGrammarBuilder creates a new grammar builder
func NewGrammarBuilder() GrammarBuilder {
	return createGrammarBuilder()
}

// Builder represents the script builder
type Builder interface {
	Create() Builder
	WithRoot(root grammars.Grammar) Builder
	WithTokens(tokens Tokens) Builder
	WithGrammars(grammars Grammars) Builder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Root() grammars.Grammar
	Tokens() Tokens
	HasGrammars() bool
	Grammars() Grammars
}

// TokensBuilder represents tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
	Fetch(hash hash.Hash) (Token, error)
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithReference(reference grammars.Token) TokenBuilder
	WithName(name string) TokenBuilder
	Now() (Token, error)
}

// Token represents a token reference
type Token interface {
	Name() string
	Reference() grammars.Token
}

// GrammarsBuilder represents the grammars builder
type GrammarsBuilder interface {
	Create() GrammarsBuilder
	WithList(list []Grammar) GrammarsBuilder
	Now() (Grammars, error)
}

// Grammars represents grammars
type Grammars interface {
	List() []Grammar
	Fetch(hash hash.Hash) (Grammar, error)
}

// GrammarBuilder represents a grammar builder
type GrammarBuilder interface {
	Create() GrammarBuilder
	WithName(name string) GrammarBuilder
	WithReference(reference grammars.Grammar) GrammarBuilder
	Now() (Grammar, error)
}

// Grammar represents a grammar reference
type Grammar interface {
	Name() string
	Reference() grammars.Grammar
}
