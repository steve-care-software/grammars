package scripts

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	"github.com/steve-care-software/grammars/infrastructure/scripts/components"
	"github.com/steve-care-software/grammars/infrastructure/scripts/tokens"
)

type grammar struct {
	component        components.Component
	token            tokens.Token
	builder          grammars.Builder
	refBuilder       references.Builder
	refTokensBuilder references.TokensBuilder
}

func createGrammar(
	component components.Component,
	token tokens.Token,
	builder grammars.Builder,
	refBuilder references.Builder,
	refTokensBuilder references.TokensBuilder,
) Grammar {
	out := grammar{
		component:        component,
		token:            token,
		builder:          builder,
		refBuilder:       refBuilder,
		refTokensBuilder: refTokensBuilder,
	}

	return &out
}

// Grammar returns the grammar's reference
func (app *grammar) Grammar() references.Reference {
	root, subRoot := app.grammarToken()
	channels, sunChannels := app.component.Channel().Channels()
	grammar, err := app.builder.Create().
		WithRoot(root.Reference()).
		WithChannels(channels).
		Now()

	if err != nil {
		panic(err)
	}

	nameTokens := []references.Token{}
	nameTokens = append(nameTokens, root)
	nameTokens = append(nameTokens, subRoot...)
	nameTokens = append(nameTokens, sunChannels...)
	scriptTokens, err := app.refTokensBuilder.Create().WithList(nameTokens).Now()
	if err != nil {
		panic(err)
	}

	ins, err := app.refBuilder.Create().WithRoot(grammar).WithTokens(scriptTokens).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) grammarToken() (references.Token, []references.Token) {
	root, subRoot := app.rootToken()
	channel, subChannel := app.channelToken()
	instruction, subInstruction := app.instructionToken()

	output := []references.Token{}
	output = append(output, root)
	output = append(output, subRoot...)
	output = append(output, channel)
	output = append(output, subChannel...)
	output = append(output, instruction)
	output = append(output, subInstruction...)

	return app.component.Token().FromLines(
		"grammar",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(root.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(channel.Reference(), app.component.Cardinality().Cardinality(0, nil)),
				app.component.Element().FromToken(instruction.Reference(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`
				@myValue;
				myValue: 45;
			`: true,
			`
				@myValue;
				-myChannel;
				-mySecondChannel [prev];

				myValue: myCompose
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				@myValue;
				-myChannel;
				-mySecondChannel [:next];

				myEverything: #myToken
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				@myValue;
				-myChannel;
				-mySecondChannel [prev:next];

				myVariable: myToken*
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
		}),
	), output
}

func (app *grammar) instructionToken() (references.Token, []references.Token) {

	valueAssignment, subValueAssignment := app.valueAssignmentToken()
	compose, subCompose := app.composeAssignmentToken()
	everything, subEverything := app.everythingAssignmentToken()
	token, subToken := app.tokenAssignmentToken()

	output := []references.Token{}
	output = append(output, valueAssignment)
	output = append(output, subValueAssignment...)
	output = append(output, compose)
	output = append(output, subCompose...)
	output = append(output, everything)
	output = append(output, subEverything...)
	output = append(output, token)
	output = append(output, subToken...)

	return app.component.Token().FromLines(
		"instruction",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(valueAssignment.Reference(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(compose.Reference(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(everything.Reference(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(token.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`myValue: 45;`: true,
			`
				myCompose: myCompose
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				myEverything: #myToken
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				myVariable: myToken*
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
		}),
	), output
}

func (app *grammar) rootToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)

	return app.component.Token().FromLines(
		"root",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(rootPrefix)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(rootSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`@myRoot;`: true,
		}),
	), output
}

func (app *grammar) channelToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	chanPrevNext, subChanPrevNext := app.channelPreviousNextToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, chanPrevNext)
	output = append(output, subChanPrevNext...)

	pMax := uint(1)
	return app.component.Token().FromLines(
		"channel",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(channelPrefix)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(chanPrevNext.Reference(), app.component.Cardinality().Cardinality(0, &pMax)),
				app.component.Element().FromValue([]byte(channelSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`-myChanToken;`:             true,
			`-myChanToken [prev];`:      true,
			`-myChanToken [:next];`:     true,
			`-myChanToken [prev:next];`: true,
		}),
	), output
}

func (app *grammar) channelPreviousNextToken() (references.Token, []references.Token) {
	channelPrevNextInside, subChannelPrevNextInside := app.channelPreviousNextInsideToken()

	output := []references.Token{}
	output = append(output, channelPrevNextInside)
	output = append(output, subChannelPrevNextInside...)

	return app.component.Token().FromLines(
		"channelPreviousNext",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(channelPrevNextPrefix)),
				app.component.Element().FromToken(channelPrevNextInside.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(channelPrevNextSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`[prev]`:      true,
			`[:next]`:     true,
			`[prev:next]`: true,
		}),
	), output
}

func (app *grammar) channelPreviousNextInsideToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)

	return app.component.Token().FromLines(
		"channelPreviousNextInside",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(channelPrevNextDelimiter)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(channelPrevNextDelimiter)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`prev`:      true,
			`:next`:     true,
			`prev:next`: true,
		}),
	), output
}

func (app *grammar) tokenAssignmentToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	blockToken, subBlockToken := app.blockToken()
	suite, subSuite := app.suiteToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, blockToken)
	output = append(output, subBlockToken...)
	output = append(output, suite)
	output = append(output, subSuite...)

	return app.component.Token().FromLines(
		"tokenAssignment",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)),
				app.component.Element().FromToken(blockToken.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suite.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`
				myVariable: myToken*
				---
					valid	: firstValidCompose
							& secondValidCompose
							;
				;
			`: true,
			`
				myVariable	: myToken*
							| mySecond+
							| myThird?
							| fourth[2] fifth[0,] sixth[1,] seventh[0,234]
				---
					valid	: firstValidCompose
							& secondValidCompose
							;

					invalid	: myComposeToken;
				;
			`: true,
		}),
	), output
}

func (app *grammar) valueAssignmentToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	anyByte := app.token.AnyByte()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, anyByte)

	return app.component.Token().FromLines(
		"valueAssignment",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)),
				app.component.Element().FromToken(anyByte.Reference(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`
				myValue: 45;
			`: true,
		}),
	), output
}

func (app *grammar) suiteToken() (references.Token, []references.Token) {
	suitePrefixCons := app.component.Token().AllCharacters("suitePrefixConst", suitePrefix)
	suiteValid, subSuiteValid := app.suiteValidToken()
	suiteInvalid, subSuiteInvalid := app.suiteInvalidToken()

	output := []references.Token{}
	output = append(output, suitePrefixCons)
	output = append(output, suiteValid)
	output = append(output, subSuiteValid...)
	output = append(output, suiteInvalid)
	output = append(output, subSuiteInvalid...)

	pMax := uint(1)
	return app.component.Token().FromLines(
		"suite",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(suitePrefixCons.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suiteValid.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suiteInvalid.Reference(), app.component.Cardinality().Cardinality(0, &pMax)),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(suitePrefixCons.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suiteValid.Reference(), app.component.Cardinality().Cardinality(0, &pMax)),
				app.component.Element().FromToken(suiteInvalid.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`
				---
			`: false,
			`
				---
				valid	: firstValidCompose
						& secondValidCompose
						;
			`: true,
			`
				---
				invalid : myComposeToken
					  	& mySecondComposeToken
						& myThirdComposeToken
					  	;
			`: true,
			`
				---
				valid	: firstValidCompose
						& secondValidCompose
						;

				invalid	: myComposeToken;
			`: true,
			`
				---
				valid	: firstValidCompose;

				invalid : myComposeToken
					  	& mySecondComposeToken
						& myThirdComposeToken
					  	;
			`: true,
		}),
	), output
}

func (app *grammar) suiteInvalidToken() (references.Token, []references.Token) {
	invalidCons := app.component.Token().AllCharacters("invalidConst", invalidSuiteName)
	suiteBlock, subSuiteBlock := app.suiteBlockToken()

	output := []references.Token{}
	output = append(output, invalidCons)
	output = append(output, suiteBlock)
	output = append(output, subSuiteBlock...)

	return app.component.Token().FromLines(
		"suiteInvalid",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(invalidCons.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suiteBlock.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`invalid: myComposeToken;`: true,
			`
				invalid : myComposeToken
					  	& mySecondComposeToken
						& myThirdComposeToken
					  	;
			`: true,
		}),
	), output
}

func (app *grammar) suiteValidToken() (references.Token, []references.Token) {
	validCons := app.component.Token().AllCharacters("validConst", validSuiteName)
	suiteBlock, subSuiteBlock := app.suiteBlockToken()

	output := []references.Token{}
	output = append(output, validCons)
	output = append(output, suiteBlock)
	output = append(output, subSuiteBlock...)

	return app.component.Token().FromLines(
		"suiteValid",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(validCons.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suiteBlock.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`valid: myComposeToken;`: true,
			`
				valid : myComposeToken
					  & mySecondComposeToken
					  ;
			`: true,
		}),
	), output
}

func (app *grammar) suiteBlockToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	delimiterThenSuite, subDelimiterThenSuite := app.delimiterThenSuiteElementToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, delimiterThenSuite)
	output = append(output, subDelimiterThenSuite...)

	return app.component.Token().FromLines(
		"suiteBlock",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(assignmentSign)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(delimiterThenSuite.Reference(), app.component.Cardinality().Cardinality(0, nil)),
				app.component.Element().FromValue([]byte(suiteSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`: myComposeToken;`:                        true,
			`: myComposeToken & mySecondComposeToken;`: true,
		}),
	), output
}

func (app *grammar) delimiterThenSuiteElementToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	return app.component.Token().FromLines(
		"delimiterThenSuiteElement",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(suiteDelimiter)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`& myComposeToken`: true,
		}),
	), append(subVariableName, variableName)
}

func (app *grammar) blockToken() (references.Token, []references.Token) {
	line, subLine := app.lineToken()
	delimiterThenLine, subDelimiterThenLine := app.delimiterThenLineToken()

	output := []references.Token{}
	output = append(output, line)
	output = append(output, subLine...)
	output = append(output, delimiterThenLine)
	output = append(output, subDelimiterThenLine...)

	return app.component.Token().FromLines(
		"block",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(line.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(delimiterThenLine.Reference(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`myToken*`: true,
			`
				myToken*
				| mySecond+
				| myThird?
				| fourth[2] fifth[0,] sixth[1,] seventh[0,234]
			`: true,
		}),
	), output
}

func (app *grammar) delimiterThenLineToken() (references.Token, []references.Token) {
	lineToken, subLineToken := app.lineToken()
	return app.component.Token().FromLines(
		"delimiterThenLine",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(lineDelimiter)),
				app.component.Element().FromToken(lineToken.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`|myToken* mySecond+ myThird? fourth[2] fifth[0,] sixth[1,] seventh[0,234]`: true,
		}),
	), append(subLineToken, lineToken)
}

func (app *grammar) lineToken() (references.Token, []references.Token) {
	element, subElement := app.elementToken()
	return app.component.Token().FromLines(
		"line",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(element.Reference(), app.component.Cardinality().Cardinality(1, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`myToken* mySecond+ myThird? fourth[2] fifth[0,] sixth[1,] seventh[0,234]`: true,
		}),
	), append(subElement, element)
}

func (app *grammar) elementToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	cardinality, subCardinality := app.cardinalityToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, cardinality)
	output = append(output, subCardinality...)

	return app.component.Token().FromLines(
		"element",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(cardinality.Reference(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`myToken*`:       true,
			`myToken+`:       true,
			`myToken?`:       true,
			`myToken[2]`:     true,
			`myToken[234]`:   true,
			`myToken[0,]`:    true,
			`myToken[1,]`:    true,
			`myToken[0,234]`: true,
		}),
	), output
}

func (app *grammar) cardinalityToken() (references.Token, []references.Token) {
	anyNumber := app.token.AnyNumber()
	return app.component.Token().FromLines(
			"cardinality",
			[]grammars.Line{
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromValue([]byte(cardinalitySingleOptional)),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromValue([]byte(cardinalityMultipleMandatory)),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromValue([]byte(cardinalityMultipleOptional)),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromValue([]byte(cardinalityPrefix)),
					app.component.Element().FromToken(anyNumber.Reference(), app.component.Cardinality().Cardinality(1, nil)),
					app.component.Element().FromValue([]byte(cardinalitySuffix)),
				}),
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromValue([]byte(cardinalityPrefix)),
					app.component.Element().FromToken(anyNumber.Reference(), app.component.Cardinality().Cardinality(1, nil)),
					app.component.Element().FromValue([]byte(cardinalitySeparator)),
					app.component.Element().FromToken(anyNumber.Reference(), app.component.Cardinality().Cardinality(0, nil)),
					app.component.Element().FromValue([]byte(cardinalitySuffix)),
				}),
			},
			app.component.Suite().Suites(map[string]bool{
				`*`:       true,
				`+`:       true,
				`?`:       true,
				`[2]`:     true,
				`[234]`:   true,
				`[0,]`:    true,
				`[1,]`:    true,
				`[0,234]`: true,
			}),
		), []references.Token{
			anyNumber,
		}
}

func (app *grammar) composeAssignmentToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	composeToken, subComposeToken := app.composeToken()
	suite, subSuite := app.suiteToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, composeToken)
	output = append(output, subComposeToken...)
	output = append(output, suite)
	output = append(output, subSuite...)

	return app.component.Token().FromLines(
		"composeAssignment",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)),
				app.component.Element().FromToken(composeToken.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suite.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`
				myCompose: myCompose
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
			`
				myCompose: myCompose|45
					---
					valid: myValidCompose;
				;
			`: true,
		}),
	), output
}

func (app *grammar) composeToken() (references.Token, []references.Token) {
	variableName, subVariableName := app.token.VariableName()
	separator, subSeparator := app.separatorAmountOfComposeToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableName...)
	output = append(output, separator)
	output = append(output, subSeparator...)

	return app.component.Token().FromLines(
		"compose",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(separator.Reference(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`myCompose`:     true,
			`myCompose|4`:   true,
			`myCompose|234`: true,
		}),
	), output
}

func (app *grammar) separatorAmountOfComposeToken() (references.Token, []references.Token) {
	anyNumber := app.token.AnyNumber()
	return app.component.Token().FromLines(
			"composeWithAmount",
			[]grammars.Line{
				app.component.Line().FromElements([]grammars.Element{
					app.component.Element().FromValue([]byte(amountSeparator)),
					app.component.Element().FromToken(anyNumber.Reference(), app.component.Cardinality().Cardinality(1, nil)),
				}),
			},
			app.component.Suite().Suites(map[string]bool{
				`|4`:   true,
				`|234`: true,
			}),
		), []references.Token{
			anyNumber,
		}
}

func (app *grammar) everythingAssignmentToken() (references.Token, []references.Token) {
	variableName, subVariableNames := app.token.VariableName()
	everything, subEverything := app.everythingToken()
	suite, subSuite := app.suiteToken()

	output := []references.Token{}
	output = append(output, variableName)
	output = append(output, subVariableNames...)
	output = append(output, everything)
	output = append(output, subEverything...)
	output = append(output, suite)
	output = append(output, subSuite...)

	return app.component.Token().FromLines(
		"everythingAssignment",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(assignmentSign)),
				app.component.Element().FromToken(everything.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(suite.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(blockSuffix)),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			`
				myEverything: #myToken
					---
					valid: myValidCompose;
					invalid : myInvalidCompose
							& mySecondInvalidCompose
							;
				;
			`: true,
		}),
	), output
}

func (app *grammar) everythingToken() (references.Token, []references.Token) {
	evWithEscapeToken, namesEvWithEscapeToken := app.everythingWithEscapeToken()
	evWithoutEscapeToken, namesEvWithoutEscapeToken := app.everythingWithoutEscapeToken()

	output := []references.Token{}
	output = append(output, evWithEscapeToken)
	output = append(output, namesEvWithEscapeToken...)
	output = append(output, evWithoutEscapeToken)
	output = append(output, namesEvWithoutEscapeToken...)

	return app.component.Token().FromLines(
		"everything",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(evWithEscapeToken.Reference(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(evWithoutEscapeToken.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			"#myToken":          true,
			"#myToken!myEscape": true,
		}),
	), output
}

func (app *grammar) everythingWithEscapeToken() (references.Token, []references.Token) {
	variableName, namesVariasbleName := app.token.VariableName()
	return app.component.Token().FromLines(
		"everythingWithEscape",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(everythingPrefixSign)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
				app.component.Element().FromValue([]byte(everythingEscapePrefixSign)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			"#myToken!myEscape": true,
		}),
	), append(namesVariasbleName, variableName)
}

func (app *grammar) everythingWithoutEscapeToken() (references.Token, []references.Token) {
	variableName, namesVariasbleName := app.token.VariableName()
	return app.component.Token().FromLines(
		"everythingWithoutEscape",
		[]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromValue([]byte(everythingPrefixSign)),
				app.component.Element().FromToken(variableName.Reference(), app.component.Cardinality().Once()),
			}),
		},
		app.component.Suite().Suites(map[string]bool{
			"#myToken": true,
		}),
	), append(namesVariasbleName, variableName)
}
