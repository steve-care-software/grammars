package channels

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	script_cardinalities "github.com/steve-care-software/grammars/infrastructure/scripts/components/cardinalities"
	script_elements "github.com/steve-care-software/grammars/infrastructure/scripts/components/elements"
	script_everythings "github.com/steve-care-software/grammars/infrastructure/scripts/components/everythings"
	script_lines "github.com/steve-care-software/grammars/infrastructure/scripts/components/lines"
	script_suites "github.com/steve-care-software/grammars/infrastructure/scripts/components/suites"
	script_tokens "github.com/steve-care-software/grammars/infrastructure/scripts/components/tokens"
)

type channel struct {
	token          script_tokens.Token
	element        script_elements.Element
	line           script_lines.Line
	suite          script_suites.Suite
	everything     script_everythings.Everything
	cardinality    script_cardinalities.Cardinality
	channelBuilder grammars.ChannelBuilder
}

func createChannel(
	token script_tokens.Token,
	element script_elements.Element,
	line script_lines.Line,
	suite script_suites.Suite,
	everything script_everythings.Everything,
	cardinality script_cardinalities.Cardinality,
	channelBuilder grammars.ChannelBuilder,
) Channel {
	out := channel{
		token:          token,
		element:        element,
		line:           line,
		suite:          suite,
		everything:     everything,
		cardinality:    cardinality,
		channelBuilder: channelBuilder,
	}

	return &out
}

// Channels returns the channels
func (app *channel) Channels() ([]grammars.Channel, []references.Token) {

	nameTokenDoubleSlash := app.token.AllCharacters("doubleSlash", "//")
	nameTokenEndOfLineSpaces := app.token.AnyElement(
		"endOfLineSpaces",
		[]grammars.Element{
			app.element.FromValue([]byte("\n")),
			app.element.FromValue([]byte("\r")),
		},
		app.suite.Suites(map[string]bool{
			"\n": true,
			"\r": true,
		}),
	)

	nameTokenSingleLine, chanSingleLineToken := app.channelFromToken(
		app.token.FromLines("singleLineComment",
			[]grammars.Line{
				app.line.FromElements([]grammars.Element{
					app.element.FromToken(
						nameTokenDoubleSlash.Reference(),
						app.cardinality.Once(),
					),
					app.element.FromEverything(
						app.everything.WithoutEscape(
							nameTokenEndOfLineSpaces.Reference(),
						),
					),
				}),
			},
			app.suite.Suites(map[string]bool{
				`// this is a comment
				`: true,
			}),
		),
	)

	nameTokenSpace, chanSpace := app.channelFromValue("space", []byte(" "))
	nameTokenTab, chanTab := app.channelFromValue("tab", []byte("\t"))
	nameTokenNewLine, chanNewLine := app.channelFromValue("newLine", []byte("\n"))
	nameTokenRetChar, chanRetChar := app.channelFromValue("retChar", []byte("\r"))
	return []grammars.Channel{
			chanSpace,
			chanTab,
			chanNewLine,
			chanRetChar,
			chanSingleLineToken,
		}, []references.Token{
			nameTokenDoubleSlash,
			nameTokenEndOfLineSpaces,
			nameTokenSpace,
			nameTokenTab,
			nameTokenNewLine,
			nameTokenRetChar,
			nameTokenSingleLine,
		}
}

func (app *channel) channelFromValue(name string, value []byte) (references.Token, grammars.Channel) {
	return app.channelFromToken(
		app.token.FromLines(
			name,
			[]grammars.Line{
				app.line.FromElements([]grammars.Element{
					app.element.FromValue(value),
				}),
			},
			app.suite.Suites(map[string]bool{}),
		),
	)
}

func (app channel) channelFromToken(token references.Token) (references.Token, grammars.Channel) {
	ref := token.Reference()
	ins, err := app.channelBuilder.Create().
		WithToken(ref).
		Now()

	if err != nil {
		panic(err)
	}

	return token, ins
}
