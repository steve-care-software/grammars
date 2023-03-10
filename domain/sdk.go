package domain

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewChannelBuilder creates a new channel builder
func NewChannelBuilder() ChannelBuilder {
	hashAdapter := hash.NewAdapter()
	return createChannelBuilder(hashAdapter)
}

// NewChannelConditionBuilder creates a new chanel condition builder
func NewChannelConditionBuilder() ChannelConditionBuilder {
	return createChannelConditionBuilder()
}

// NewInstanceBuilder creates a new instance builder
func NewInstanceBuilder() InstanceBuilder {
	return createInstanceBuilder()
}

// NewEverythingBuilder creates a new everything builder
func NewEverythingBuilder() EverythingBuilder {
	hashAdapter := hash.NewAdapter()
	return createEverythingBuilder(hashAdapter)
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokenBuilder(hashAdapter)
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuiteBuilder(hashAdapter)
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

// NewCardinalityBuilder creates a new cardinality builder
func NewCardinalityBuilder() CardinalityBuilder {
	return createCardinalityBuilder()
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root Token) Builder
	WithChannels(channels []Channel) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Hash() hash.Hash
	Root() Token
	HasChannels() bool
	Channels() []Channel
}

// ChannelBuilder represents a channel builder
type ChannelBuilder interface {
	Create() ChannelBuilder
	WithToken(token Token) ChannelBuilder
	WithCondition(condition ChannelCondition) ChannelBuilder
	Now() (Channel, error)
}

// Channel represents a channel
type Channel interface {
	Hash() hash.Hash
	Token() Token
	HasCondition() bool
	Condition() ChannelCondition
}

// ChannelConditionBuilder represents a channel condition builder
type ChannelConditionBuilder interface {
	Create() ChannelConditionBuilder
	WithPrevious(previous Token) ChannelConditionBuilder
	WithNext(next Token) ChannelConditionBuilder
	Now() (ChannelCondition, error)
}

// ChannelCondition represents a channel condition
type ChannelCondition interface {
	Hash() hash.Hash
	HasPrevious() bool
	Previous() Token
	HasNext() bool
	Next() Token
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithLines(lines []Line) TokenBuilder
	WithSuites(suites []Suite) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Lines() []Line
	HasSuites() bool
	Suites() []Suite
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithValid(valid []byte) SuiteBuilder
	WithInvalid(invalid []byte) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a test suite
type Suite interface {
	Hash() hash.Hash
	IsValid() bool
	Content() []byte
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithElements(elements []Element) LineBuilder
	Now() (Line, error)
}

// Line represents a line of elements
type Line interface {
	Hash() hash.Hash
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithCardinality(cardinality Cardinality) ElementBuilder
	WithValue(value []byte) ElementBuilder
	WithGrammar(grammar Grammar) ElementBuilder
	WithInstance(instance Instance) ElementBuilder
	WithRecursive(recursive string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Content() ElementContent
	Cardinality() Cardinality
}

// ElementContent represents an element content
type ElementContent interface {
	Hash() hash.Hash
	IsValue() bool
	Value() []byte
	IsGrammar() bool
	Grammar() Grammar
	IsInstance() bool
	Instance() Instance
	IsRecursive() bool
	Recursive() string
}

// InstanceBuilder represents an instance builder
type InstanceBuilder interface {
	Create() InstanceBuilder
	WithToken(token Token) InstanceBuilder
	WithEverything(everything Everything) InstanceBuilder
	Now() (Instance, error)
}

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
	IsToken() bool
	Token() Token
	IsEverything() bool
	Everything() Everything
}

// EverythingBuilder represents an everything builder
type EverythingBuilder interface {
	Create() EverythingBuilder
	WithException(exception Token) EverythingBuilder
	WithEscape(escape Token) EverythingBuilder
	Now() (Everything, error)
}

// Everything represents an everything except
type Everything interface {
	Hash() hash.Hash
	Exception() Token
	HasEscape() bool
	Escape() Token
}

// CardinalityBuilder represents a cardinality builder
type CardinalityBuilder interface {
	Create() CardinalityBuilder
	WithMin(min uint) CardinalityBuilder
	WithMax(max uint) CardinalityBuilder
	Now() (Cardinality, error)
}

// Cardinality represents a cardinality
type Cardinality interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
