package scripts

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components"
	"github.com/steve-care-software/grammars/infrastructure/scripts/tokens"
)

const grammarTokenName = "grammar"

const assignmentSign = ":"
const everythingPrefixSign = "#"
const everythingEscapePrefixSign = "!"
const amountSeparator = "|"
const cardinalitySingleOptional = "?"
const cardinalityMultipleMandatory = "+"
const cardinalityMultipleOptional = "*"
const cardinalityPrefix = "["
const cardinalitySuffix = "]"
const cardinalitySeparator = ","
const lineDelimiter = "|"
const blockSuffix = ";"
const validSuiteName = "valid"
const invalidSuiteName = "invalid"
const suiteDelimiter = "&"
const suitePrefix = "---"
const suiteSuffix = ";"
const channelPrefix = "-"
const channelSuffix = ";"
const channelPrevNextPrefix = "["
const channelPrevNextSuffix = "]"
const channelPrevNextDelimiter = ":"
const rootPrefix = "@"
const rootSuffix = ";"
const instructionSuffix = ";"
const externalTokenPrefix = "{"
const externalTokenSuffix = "{"

// NewGrammar creates a new grammar instance
func NewGrammar() Grammar {
	component := components.NewComponent()
	token := tokens.NewToken()
	builder := grammars.NewBuilder()
	refBuilder := references.NewBuilder()
	refTokensBuilder := references.NewTokensBuilder()
	return createGrammar(
		component,
		token,
		builder,
		refBuilder,
		refTokensBuilder,
	)
}

// Grammar represents the grammar
type Grammar interface {
	Grammar() references.Reference
}
