package tokens

import (
	"github.com/steve-care-software/grammars/domain/references"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components"
)

const byteLength = 256

// NewToken creates a new token instance
func NewToken() Token {
	component := components.NewComponent()
	return createToken(
		component,
	)
}

// Token represents the token component
type Token interface {
	VariableName() (references.Token, []references.Token)
	Sha512Hex() (references.Token, []references.Token)
	AnyHexCharacter() (references.Token, []references.Token)
	AnyLetter() (references.Token, []references.Token)
	AnyByte() references.Token
	AnyNumber() references.Token
	AToFUpperCaseLetters() references.Token
	AToFLowerCaseLetters() references.Token
	UpperCaseLetters() references.Token
	LowerCaseLetters() references.Token
}
