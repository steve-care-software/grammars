package tokens

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	script_elements "github.com/steve-care-software/grammars/infrastructure/scripts/components/elements"
	script_lines "github.com/steve-care-software/grammars/infrastructure/scripts/components/lines"
	script_suites "github.com/steve-care-software/grammars/infrastructure/scripts/components/suites"
)

type token struct {
	suite           script_suites.Suite
	element         script_elements.Element
	line            script_lines.Line
	tokenBuilder    grammars.TokenBuilder
	refTokenBuilder references.TokenBuilder
}

func createToken(
	suite script_suites.Suite,
	element script_elements.Element,
	line script_lines.Line,
	tokenBuilder grammars.TokenBuilder,
	refTokenBuilder references.TokenBuilder,
) Token {
	out := token{
		suite:           suite,
		element:         element,
		line:            line,
		tokenBuilder:    tokenBuilder,
		refTokenBuilder: refTokenBuilder,
	}

	return &out
}

// AllCharacters returns the all character token
func (app *token) AllCharacters(tokenName string, values string) references.Token {
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.element.FromValue([]byte{byte(oneValue)})
		elementsList = append(elementsList, element)
	}

	return app.allElementsToken(
		tokenName,
		elementsList,
		app.suite.Suites(map[string]bool{
			values: true,
		}),
	)
}

// AnyCharacter returns the any character token
func (app *token) AnyCharacter(tokenName string, values string) references.Token {
	suitesData := map[string]bool{}
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.element.FromValue([]byte{byte(oneValue)})
		elementsList = append(elementsList, element)
		suitesData[string(oneValue)] = true
	}

	return app.AnyElement(
		tokenName,
		elementsList,
		app.suite.Suites(suitesData),
	)
}

// AnyElement returns an any element token
func (app *token) AnyElement(tokenName string, elementsList []grammars.Element, suites []grammars.Suite) references.Token {
	linesList := []grammars.Line{}
	for _, oneElement := range elementsList {
		linesList = append(
			linesList,
			app.line.FromElements([]grammars.Element{
				oneElement,
			}),
		)
	}

	return app.FromLines(
		tokenName,
		linesList,
		suites,
	)
}

// FromLines returns a token from lines
func (app *token) FromLines(name string, lines []grammars.Line, suites []grammars.Suite) references.Token {
	ref, err := app.tokenBuilder.Create().
		WithLines(lines).
		WithSuites(suites).
		Now()

	if err != nil {
		panic(err)
	}

	ins, err := app.refTokenBuilder.Create().
		WithReference(ref).
		WithName(name).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *token) allCharacterToken(tokenName string, values string, suites []grammars.Suite) references.Token {
	return app.allElementsToken(
		tokenName,
		[]grammars.Element{
			app.element.FromValue([]byte(values)),
		},
		suites,
	)
}

func (app *token) allElementsToken(tokenName string, elementsList []grammars.Element, suites []grammars.Suite) references.Token {
	return app.FromLines(
		tokenName,
		[]grammars.Line{
			app.line.FromElements(elementsList),
		},
		suites,
	)
}
