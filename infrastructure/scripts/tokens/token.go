package tokens

import (
	"fmt"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components"
)

type token struct {
	component components.Component
}

func createToken(
	component components.Component,
) Token {
	out := token{
		component: component,
	}

	return &out
}

// VariableName returns the variable name token
func (app *token) VariableName() (references.Token, []references.Token) {
	nameLowerCaseLetter := app.LowerCaseLetters()
	nameAnyLetter, subNameAnyLetter := app.AnyLetter()

	output := []references.Token{}
	output = append(output, nameLowerCaseLetter)
	output = append(output, nameAnyLetter)
	output = append(output, subNameAnyLetter...)

	return app.component.Token().FromLines(
		"variableName",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(nameLowerCaseLetter.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(nameAnyLetter.Reference(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			"m":          true,
			"myVariable": true,
			"MyVariable": false,
			"0Variable":  false,
		}),
	), output
}

// Sha512Hex returns the sha512 token
func (app *token) Sha512Hex() (references.Token, []references.Token) {
	amount := uint(128)
	nameAnyHexCharacter, subNameTokenAnyHexCharacterList := app.AnyHexCharacter()
	return app.component.Token().FromLines(
		"sha512Hex",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(nameAnyHexCharacter.Reference(), app.component.Cardinality().Cardinality(amount, &amount)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			"cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c": true,
			"cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d9460":   false,
			"gf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c": false,
		}),
	), append(subNameTokenAnyHexCharacterList, nameAnyHexCharacter)
}

// AnyHexCharacter returns the [a-fA-Z0-9] token
func (app *token) AnyHexCharacter() (references.Token, []references.Token) {

	nameAFLower := app.AToFLowerCaseLetters()
	nameAFUpper := app.AToFUpperCaseLetters()
	nameAnyNumber := app.AnyNumber()

	return app.component.Token().FromLines(
			"anyHexChar",
			[]grammars.Line{
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromToken(nameAFLower.Reference(), app.component.Cardinality().Once()),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromToken(nameAFUpper.Reference(), app.component.Cardinality().Once()),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromToken(nameAnyNumber.Reference(), app.component.Cardinality().Once()),
				}),
			},
			app.component.Suite().Suites(map[string]bool{
				"a": true,
				"b": true,
				"c": true,
				"d": true,
				"e": true,
				"f": true,
				"g": false,
				"h": false,
				"i": false,
				"j": false,
				"k": false,
				"l": false,
				"m": false,
				"n": false,
				"o": false,
				"p": false,
				"q": false,
				"r": false,
				"s": false,
				"t": false,
				"u": false,
				"v": false,
				"w": false,
				"x": false,
				"y": false,
				"z": false,
				"A": true,
				"B": true,
				"C": true,
				"D": true,
				"E": true,
				"F": true,
				"G": false,
				"H": false,
				"I": false,
				"J": false,
				"K": false,
				"L": false,
				"M": false,
				"N": false,
				"O": false,
				"P": false,
				"Q": false,
				"R": false,
				"S": false,
				"T": false,
				"U": false,
				"V": false,
				"W": false,
				"X": false,
				"Y": false,
				"Z": false,
				"0": true,
				"1": true,
				"2": true,
				"3": true,
				"4": true,
				"5": true,
				"6": true,
				"7": true,
				"8": true,
				"9": true,
			}),
		), []references.Token{
			nameAFLower,
			nameAFUpper,
			nameAnyNumber,
		}
}

// AnyLetter returns the [a-zA-Z] token
func (app *token) AnyLetter() (references.Token, []references.Token) {
	nameUpperCaseLetters := app.UpperCaseLetters()
	nameLowerCaseLetters := app.LowerCaseLetters()
	return app.component.Token().FromLines(
			"anyLetter",
			[]grammars.Line{
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromToken(nameUpperCaseLetters.Reference(), app.component.Cardinality().Once()),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromToken(nameLowerCaseLetters.Reference(), app.component.Cardinality().Once()),
				}),
			},
			app.component.Suite().Suites(map[string]bool{
				"a": true,
				"b": true,
				"c": true,
				"d": true,
				"e": true,
				"f": true,
				"g": true,
				"h": true,
				"i": true,
				"j": true,
				"k": true,
				"l": true,
				"m": true,
				"n": true,
				"o": true,
				"p": true,
				"q": true,
				"r": true,
				"s": true,
				"t": true,
				"u": true,
				"v": true,
				"w": true,
				"x": true,
				"y": true,
				"z": true,
				"A": true,
				"B": true,
				"C": true,
				"D": true,
				"E": true,
				"F": true,
				"G": true,
				"H": true,
				"I": true,
				"J": true,
				"K": true,
				"L": true,
				"M": true,
				"N": true,
				"O": true,
				"P": true,
				"Q": true,
				"R": true,
				"S": true,
				"T": true,
				"U": true,
				"V": true,
				"W": true,
				"X": true,
				"Y": true,
				"Z": true,
				"0": false,
			}),
		), []references.Token{
			nameUpperCaseLetters,
			nameLowerCaseLetters,
		}
}

// AnyByte returns the any byte token
func (app *token) AnyByte() references.Token {
	validData := map[string]bool{}
	tokenName := "anyByte"
	elementsList := []grammars.Element{}
	for i := 0; i < byteLength; i++ {
		element := app.component.Element().FromValue([]byte{byte(i)})
		elementsList = append(elementsList, element)
		validData[fmt.Sprintf("%d", i)] = true
	}

	return app.component.Token().AnyElement(tokenName, elementsList, app.component.Suite().Suites(validData))
}

// AnyNumber returns the [0-9] token
func (app *token) AnyNumber() references.Token {
	characters := "0123456789"
	return app.component.Token().AnyCharacter("anyNumber", characters)
}

// AToFUpperCaseLetters returns the [A-F] token
func (app *token) AToFUpperCaseLetters() references.Token {
	characters := "ABCDEF"
	return app.component.Token().AnyCharacter("aToFUpperCaseLetters", characters)
}

// AToFLowerCaseLetters returns the [a-f] token
func (app *token) AToFLowerCaseLetters() references.Token {
	characters := "abcdef"
	return app.component.Token().AnyCharacter("aToFLowerCaseLetter", characters)
}

// UpperCaseLetters returns the uppercase letter token
func (app *token) UpperCaseLetters() references.Token {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return app.component.Token().AnyCharacter("uppercaseLetter", characters)
}

// LowerCaseLetters returns the lowercase letters token
func (app *token) LowerCaseLetters() references.Token {
	characters := "abcdefghijklmnopqrstuvwxyz"
	return app.component.Token().AnyCharacter("lowerCaseLetter", characters)
}
