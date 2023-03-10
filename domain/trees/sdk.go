package trees

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewTreeBuilder creates a new tree builder instance
func NewTreeBuilder() TreeBuilder {
	hashAdapter := hash.NewAdapter()
	return createTreeBuilder(hashAdapter)
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokenBuilder(hashAdapter)
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	hashAdapter := hash.NewAdapter()
	return createLineBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createValueBuilder(hashAdapter)
}

// Builder represents a trees builder
type Builder interface {
	Create() Builder
	WithList(list []Tree) Builder
	Now() (Trees, error)
}

// Trees represents a trees
type Trees interface {
	Bytes(includeChannels bool) []byte
	Hash() hash.Hash
	List() []Tree
}

// TreeBuilder represents a tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithGrammar(grammar grammars.Token) TreeBuilder
	WithToken(token Token) TreeBuilder
	WithSuffix(suffix Trees) TreeBuilder
	WithRemaining(remaining []byte) TreeBuilder
	Now() (Tree, error)
}

// Tree represents a tree
type Tree interface {
	Fetch(hash hash.Hash, elementIndex uint) (Tree, Element, error)
	Bytes(includeChannels bool) []byte
	Hash() hash.Hash
	Grammar() grammars.Token
	Token() Token
	HasSuffix() bool
	Suffix() Trees
	HasRemaining() bool
	Remaining() []byte
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithLines(lines []Line) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Lines() []Line
	HasSuccessful() bool
	Successful() Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithIndex(index uint) LineBuilder
	WithGrammar(grammar grammars.Line) LineBuilder
	WithElements(elements []Element) LineBuilder
	IsReverse() LineBuilder
	Now() (Line, error)
}

// Line represents a line of elements
type Line interface {
	Fetch(hash hash.Hash) (Element, error)
	IsSuccessful() bool
	Hash() hash.Hash
	Index() uint
	Grammar() grammars.Line
	IsReverse() bool
	HasElements() bool
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGrammar(grammar grammars.Element) ElementBuilder
	WithContents(contents []Content) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Fetch(hash hash.Hash, elementIndex uint) (Tree, Element, error)
	Bytes(includeChannels bool) []byte
	Hash() hash.Hash
	IsSuccessful() bool
	Contents() []Content
	Amount() uint
	HasGrammar() bool
	Grammar() grammars.Element
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithValue(value Value) ContentBuilder
	WithTree(tree Tree) ContentBuilder
	Now() (Content, error)
}

// Content represents an element token
type Content interface {
	Bytes(includeChannels bool) []byte
	Hash() hash.Hash
	IsValue() bool
	Value() Value
	IsTree() bool
	Tree() Tree
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithContent(content []byte) ValueBuilder
	WithPrefix(prefix Trees) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Hash() hash.Hash
	Content() []byte
	HasPrefix() bool
	Prefix() Trees
}
