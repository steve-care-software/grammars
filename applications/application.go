package applications

import (
	"bytes"
	"errors"
	"fmt"

	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	"github.com/steve-care-software/grammars/domain/references/coverages"
	"github.com/steve-care-software/grammars/domain/trees"
)

type application struct {
	grammarTokenBuilder       grammars.TokenBuilder
	treesBuilder              trees.Builder
	treeBuilder               trees.TreeBuilder
	treeTokenBuilder          trees.TokenBuilder
	treeLineBuilder           trees.LineBuilder
	treeElementBuilder        trees.ElementBuilder
	treeContentBuilder        trees.ContentBuilder
	treeValueBuilder          trees.ValueBuilder
	coveragesBuilder          coverages.Builder
	coverageBuilder           coverages.CoverageBuilder
	coverageExecutionsBuilder coverages.ExecutionsBuilder
	coverageExecutionBuilder  coverages.ExecutionBuilder
	coverageResultBuilder     coverages.ResultBuilder
}

func createApplication(
	grammarTokenBuilder grammars.TokenBuilder,
	treesBuilder trees.Builder,
	treeBuilder trees.TreeBuilder,
	treeTokenBuilder trees.TokenBuilder,
	treeLineBuilder trees.LineBuilder,
	treeElementBuilder trees.ElementBuilder,
	treeContentBuilder trees.ContentBuilder,
	treeValueBuilder trees.ValueBuilder,
	coveragesBuilder coverages.Builder,
	coverageBuilder coverages.CoverageBuilder,
	coverageExecutionsBuilder coverages.ExecutionsBuilder,
	coverageExecutionBuilder coverages.ExecutionBuilder,
	coverageResultBuilder coverages.ResultBuilder,
) Application {
	out := application{
		grammarTokenBuilder:       grammarTokenBuilder,
		treesBuilder:              treesBuilder,
		treeBuilder:               treeBuilder,
		treeTokenBuilder:          treeTokenBuilder,
		treeLineBuilder:           treeLineBuilder,
		treeElementBuilder:        treeElementBuilder,
		treeContentBuilder:        treeContentBuilder,
		treeValueBuilder:          treeValueBuilder,
		coveragesBuilder:          coveragesBuilder,
		coverageBuilder:           coverageBuilder,
		coverageExecutionsBuilder: coverageExecutionsBuilder,
		coverageExecutionBuilder:  coverageExecutionBuilder,
		coverageResultBuilder:     coverageResultBuilder,
	}

	return &out
}

// Execute executes grammar on data
func (app *application) Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error) {
	return app.grammar(grammar, false, []byte{}, values)
}

// Coverages returns the coverages of a grammar
func (app *application) Coverages(reference references.Reference) (coverages.Coverages, error) {
	grammar := reference.Root()
	return app.coverages(reference, grammar)
}

func (app *application) coverages(reference references.Reference, grammar grammars.Grammar) (coverages.Coverages, error) {
	root := grammar.Root()
	channels := grammar.Channels()
	skip := map[string]bool{}
	rootCoverages, err := app.coveragesToken(reference, root, channels, &skip)
	if err != nil {
		return nil, err
	}

	list := []coverages.Coverage{}
	if grammar.HasChannels() {
		channels := grammar.Channels()
		for _, oneChannel := range channels {
			token := oneChannel.Token()
			coverages, err := app.coveragesToken(reference, token, nil, &skip)
			if err != nil {
				return nil, err
			}

			if coverages != nil {
				list = append(list, coverages.List()...)
			}
		}
	}

	if rootCoverages != nil {
		list = append(list, rootCoverages.List()...)
	}

	if len(list) <= 0 {
		return nil, nil
	}

	return app.coveragesBuilder.Create().WithList(list).Now()
}

func (app *application) coveragesToken(reference references.Reference, token grammars.Token, channels []grammars.Channel, pSkip *map[string]bool) (coverages.Coverages, error) {
	tokenHashStr := token.Hash().String()
	skip := *pSkip
	if _, ok := skip[tokenHashStr]; ok {
		return nil, nil
	}

	skip[tokenHashStr] = true
	pSkip = &skip
	executionsList := []coverages.Execution{}
	if token.HasSuites() {
		suites := token.Suites()
		for _, oneSuite := range suites {
			execution, err := app.coverageTokenSuite(reference, token, channels, oneSuite)
			if err != nil {
				return nil, err
			}

			if execution == nil {
				continue
			}

			executionsList = append(executionsList, execution)
		}
	}

	list := []coverages.Coverage{}
	lines := token.Lines()
	for _, oneLine := range lines {
		elements := oneLine.Elements()
		for _, oneElement := range elements {
			content := oneElement.Content()
			if content.IsGrammar() {
				grammar := content.Grammar()
				coverages, err := app.coverages(reference, grammar)
				if err != nil {
					return nil, err
				}

				if coverages != nil {
					list = append(list, coverages.List()...)
				}
			}

			if content.IsInstance() {
				instance := content.Instance()
				if instance.IsToken() {
					token := instance.Token()
					coverages, err := app.coveragesToken(reference, token, channels, pSkip)
					if err != nil {
						return nil, err
					}

					if coverages != nil {
						list = append(list, coverages.List()...)
					}
				}

				if instance.IsEverything() {
					everything := instance.Everything()
					exception := everything.Exception()
					coverages, err := app.coveragesToken(reference, exception, channels, pSkip)
					if err != nil {
						return nil, err
					}

					if coverages != nil {
						list = append(list, coverages.List()...)
					}

					if everything.HasEscape() {
						escape := everything.Escape()
						coverages, err := app.coveragesToken(reference, escape, channels, pSkip)
						if err != nil {
							return nil, err
						}

						if coverages != nil {
							list = append(list, coverages.List()...)
						}
					}
				}
			}
		}
	}

	if len(executionsList) > 0 {
		executions, err := app.coverageExecutionsBuilder.Create().WithList(executionsList).Now()
		if err != nil {
			return nil, err
		}

		referenceToken, err := reference.Tokens().Fetch(token.Hash())
		if err != nil {
			return nil, err
		}

		coverage, err := app.coverageBuilder.Create().WithToken(referenceToken).WithExecutions(executions).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, coverage)
	}

	if len(list) <= 0 {
		return nil, nil
	}

	return app.coveragesBuilder.Create().WithList(list).Now()
}

func (app *application) coverageTokenSuite(reference references.Reference, token grammars.Token, channels []grammars.Channel, suite grammars.Suite) (coverages.Execution, error) {
	input := suite.Content()
	tree, _, err := app.token(token, map[string]*stack{}, nil, channels, false, []byte{}, input)
	resultBuilder := app.coverageResultBuilder.Create()
	if tree != nil {
		resultBuilder.WithTree(tree)
	}

	if err != nil {
		resultBuilder.WithError(err.Error())
	}

	result, err := resultBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.coverageExecutionBuilder.Create().
		WithExpectation(suite).
		WithResult(result).
		Now()
}

func (app *application) findElements(reference references.Reference, grammar grammars.Grammar, pElements *map[string]map[uint]map[uint]string) error {
	elements := *pElements
	root := grammar.Root()
	err := app.findElementsFromToken(reference, root, &elements)
	if err != nil {
		return err
	}

	if grammar.HasChannels() {
		channels := grammar.Channels()
		for _, oneChannel := range channels {
			token := oneChannel.Token()
			err := app.findElementsFromToken(reference, token, &elements)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) findElementsFromToken(reference references.Reference, token grammars.Token, pElements *map[string]map[uint]map[uint]string) error {
	elements := *pElements
	tokenHashStr := token.Hash().String()
	if _, ok := elements[tokenHashStr]; !ok {
		elements[tokenHashStr] = map[uint]map[uint]string{}
	}

	lines := token.Lines()
	for idx, oneLine := range lines {
		castedIdx := uint(idx)
		if _, ok := elements[tokenHashStr][castedIdx]; !ok {
			elements[tokenHashStr][castedIdx] = map[uint]string{}
		}

		elementsList := oneLine.Elements()
		for elementIdx, oneElement := range elementsList {
			castedElIdx := uint(elementIdx)
			elements[tokenHashStr][castedIdx][castedElIdx] = oneElement.Hash().String()

			content := oneElement.Content()
			if content.IsGrammar() {
				grammar := content.Grammar()
				err := app.findElements(reference, grammar, &elements)
				if err != nil {
					return err
				}
			}

			if content.IsInstance() {
				instance := content.Instance()
				if instance.IsToken() {
					token := instance.Token()
					err := app.findElementsFromToken(reference, token, &elements)
					if err != nil {
						return err
					}
				}

				if instance.IsEverything() {
					everything := instance.Everything()
					exception := everything.Exception()
					err := app.findElementsFromToken(reference, exception, &elements)
					if err != nil {
						return err
					}

					if everything.HasEscape() {
						escape := everything.Escape()
						err := app.findElementsFromToken(reference, escape, &elements)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	pElements = &elements
	return nil
}

func (app *application) findCoveraredElements(reference references.Reference, coverages coverages.Coverages, pCovered *map[string]map[uint]map[uint]string) error {
	list := coverages.List()
	for _, oneCoverage := range list {
		tokenName := oneCoverage.Token().Name()
		executionsList := oneCoverage.Executions().List()
		for _, oneExecution := range executionsList {
			result := oneExecution.Result()
			if !result.IsTree() {
				continue
			}

			token := result.Tree().Token()
			err := app.findCoveraredElementsFromToken(reference, tokenName, token, pCovered)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) findCoveraredElementsFromToken(reference references.Reference, tokenName string, token trees.Token, pCovered *map[string]map[uint]map[uint]string) error {
	if !token.HasSuccessful() {
		return nil
	}

	line := token.Successful()
	index := line.Index()
	elementsList := line.Elements()
	for elIdx, oneElement := range elementsList {
		if !oneElement.HasGrammar() {
			continue
		}

		grammar := oneElement.Grammar()
		if !reference.HasGrammars() {
			str := fmt.Sprintf("the grammar (hash: %s) is not name referenced", grammar.Hash().String())
			return errors.New(str)
		}

		referenceGrammar, err := reference.Grammars().Fetch(grammar.Hash())
		if err != nil {
			return err
		}

		elementName := referenceGrammar.Name()
		covered := *pCovered
		if _, ok := covered[tokenName]; !ok {
			covered[tokenName] = map[uint]map[uint]string{}
		}

		if _, ok := covered[tokenName][index]; !ok {
			covered[tokenName][index] = map[uint]string{}
		}

		castedElIdx := uint(elIdx)
		if _, ok := covered[tokenName][index][castedElIdx]; !ok {
			covered[tokenName][index][castedElIdx] = elementName
		}

		contents := oneElement.Contents()
		for _, oneContent := range contents {
			if oneContent.IsTree() {
				subToken := oneContent.Tree().Token()
				err := app.findCoveraredElementsFromToken(reference, elementName, subToken, &covered)
				if err != nil {
					return err
				}
			}
		}

		pCovered = &covered
	}

	return nil
}

func (app *application) grammar(grammar grammars.Grammar, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	root := grammar.Root()
	channels := grammar.Channels()
	tree, _, err := app.token(root, map[string]*stack{}, nil, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func (app *application) token(token grammars.Token, stackMap map[string]*stack, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, map[string]*stack, error) {
	tokenHashStr := token.Hash().String()
	if _, ok := stackMap[tokenHashStr]; !ok {
		stackMap[tokenHashStr] = &stack{
			token: token,
			lines: map[int][]byte{},
		}
	}

	tokenLines := token.Lines()
	treeToken, remaining, retStackMap, err := app.lines(token, stackMap, tokenLines, escape, channels, isReverse, prevData, currentData)
	delete(stackMap, tokenHashStr)
	if err != nil {
		return nil, nil, err
	}

	stackMap = retStackMap
	if treeToken == nil {
		str := fmt.Sprintf("there was no line discovered in the token (hash: %s) using the given data: %s", tokenHashStr, currentData)
		return nil, nil, errors.New(str)
	}

	builder := app.treeBuilder.Create().WithGrammar(token).WithToken(treeToken)
	if channels != nil {
		suffix, rem, err := app.channels(channels, prevData, remaining)
		if err == nil {
			builder.WithSuffix(suffix)
			remaining = rem
		}
	}

	if len(remaining) > 0 {
		builder.WithRemaining(remaining)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, stackMap, nil
}

func (app *application) external(external grammars.Grammar, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	treeIns, err := app.grammar(external, isReverse, prevData, currentData)
	if err != nil {
		return nil, err
	}

	lines := external.Root().Lines()
	grammarRoot, err := app.grammarTokenBuilder.Create().WithLines(lines).Now()
	if err != nil {
		return nil, err
	}

	treeToken := treeIns.Token()
	return app.treeBuilder.Create().WithGrammar(grammarRoot).WithToken(treeToken).Now()
}

func (app *application) lines(token grammars.Token, stackMap map[string]*stack, lines []grammars.Line, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Token, []byte, map[string]*stack, error) {
	tokenHashstr := token.Hash().String()
	list := []trees.Line{}
	remaining := currentData
	currentStack := stackMap

	for idx, oneLine := range lines {
		// if we already went through this line, with the same data, in the stack, skip it to avoid infinite loops:
		if _, ok := currentStack[tokenHashstr]; ok {
			if data, ok := currentStack[tokenHashstr].lines[idx]; ok {
				if bytes.Compare(remaining, data) == 0 {
					continue
				}

			}
		}

		if _, ok := currentStack[tokenHashstr]; !ok {
			currentStack[tokenHashstr] = &stack{
				token: token,
				lines: map[int][]byte{},
			}
		}

		currentStack[tokenHashstr].lines[idx] = remaining

		// if the line is in reverse:
		if isReverse {
			previousData := prevData
			contentsList := []trees.Content{}

			for {
				if len(remaining) <= 0 {
					break
				}

				if escape != nil {
					escapeTree, _, err := app.token(escape, stackMap, nil, channels, false, previousData, remaining)
					if err == nil {
						if escapeTree.Token().HasSuccessful() {
							if escapeTree.HasRemaining() {
								escapeRemaining := escapeTree.Remaining()
								treeLine, rem, _, err := app.line(currentStack, oneLine, uint(idx), escape, channels, isReverse, remaining, escapeRemaining)
								if err == nil && treeLine.IsSuccessful() {
									amount := len(escapeRemaining) - len(rem)
									values := escapeRemaining[:amount]
									for _, oneValue := range values {
										value, err := app.treeValueBuilder.Create().WithContent([]byte{oneValue}).Now()
										if err != nil {
											return nil, nil, nil, err
										}

										contentIns, err := app.treeContentBuilder.Create().WithValue(value).Now()
										if err != nil {
											return nil, nil, nil, err
										}

										contentsList = append(contentsList, contentIns)
									}

									previousData = escapeRemaining
									remaining = escapeRemaining[amount:]
								}
							}
						}
					}
				}

				_, _, _, err := app.line(currentStack, oneLine, uint(idx), escape, channels, isReverse, previousData, remaining)
				if err == nil {
					break
				}

				value, err := app.treeValueBuilder.Create().WithContent([]byte{
					remaining[0],
				}).Now()

				if err != nil {
					return nil, nil, nil, err
				}

				contentIns, err := app.treeContentBuilder.Create().WithValue(value).Now()
				if err != nil {
					return nil, nil, nil, err
				}

				contentsList = append(contentsList, contentIns)
				previousData = remaining
				remaining = remaining[1:]
			}

			elementIns, err := app.treeElementBuilder.Create().WithContents(contentsList).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			lineIns, err := app.treeLineBuilder.Create().
				WithIndex(uint(idx)).
				WithGrammar(oneLine).
				WithElements([]trees.Element{
					elementIns,
				}).
				IsReverse().
				Now()

			if err != nil {
				return nil, nil, nil, err
			}

			list = append(list, lineIns)
			break
		}

		// the line is NOT in reverse:
		lineIns, rem, retStack, err := app.line(currentStack, oneLine, uint(idx), escape, channels, isReverse, prevData, remaining)
		if err != nil {
			continue
		}

		// add the line to the list:
		list = append(list, lineIns)
		if lineIns.IsSuccessful() {
			remaining = rem
			currentStack = retStack
			break
		}
	}

	// if there is no line:
	if len(list) <= 0 {
		return nil, remaining, currentStack, nil
	}

	blockIns, err := app.treeTokenBuilder.Create().WithLines(list).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return blockIns, remaining, currentStack, nil
}

func (app *application) line(stackMap map[string]*stack, line grammars.Line, index uint, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Line, []byte, map[string]*stack, error) {
	list := []trees.Element{}
	grElements := line.Elements()
	remaining := currentData
	previousData := prevData
	currentStack := stackMap
	for _, oneElement := range grElements {
		contentsList := []trees.Content{}
		cardinality := oneElement.Cardinality()
		pMax := cardinality.Max()
		for {

			if len(remaining) <= 0 {
				break
			}

			if cardinality.HasMax() {
				amount := uint(len(contentsList))
				if amount >= *pMax {
					break
				}
			}

			contentIns, rem, retStack, err := app.element(oneElement, currentStack, escape, channels, isReverse, previousData, remaining)
			if err != nil {
				break
			}

			currentStack = retStack
			contentsList = append(contentsList, contentIns)
			previousData = remaining
			remaining = rem
		}

		min := int(cardinality.Min())
		if len(contentsList) < min {
			str := fmt.Sprintf("the expected minimum content amount (%d) was not reached (%d) and therefore the element is invalid", min, len(contentsList))
			return nil, nil, nil, errors.New(str)
		}

		if len(contentsList) > 0 {
			elementIns, err := app.treeElementBuilder.Create().WithGrammar(oneElement).WithContents(contentsList).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			list = append(list, elementIns)
		}
	}

	builder := app.treeLineBuilder.Create().
		WithIndex(index).
		WithGrammar(line)

	if len(list) > 0 {
		builder.WithElements(list)
	}

	lineIns, err := builder.Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return lineIns, remaining, currentStack, nil
}

func (app *application) element(element grammars.Element, stackMap map[string]*stack, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Content, []byte, map[string]*stack, error) {
	if len(currentData) <= 0 {
		return nil, nil, nil, errors.New("no remaining data")
	}

	content := element.Content()
	value, tree, rem, retStack, err := app.elementContent(content, stackMap, escape, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, nil, nil, err
	}

	if value == nil && tree == nil {
		return nil, nil, nil, errors.New("no value/tree found")
	}

	if tree != nil && !tree.Token().HasSuccessful() {
		return nil, nil, nil, errors.New("no successfull tree found")
	}

	contentBuilder := app.treeContentBuilder.Create()
	if value != nil {
		contentBuilder.WithValue(value)
	}

	if tree != nil {
		contentBuilder.WithTree(tree)
	}

	contentIns, err := contentBuilder.Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return contentIns, rem, retStack, nil
}

func (app *application) elementContent(content grammars.ElementContent, stackMap map[string]*stack, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Value, trees.Tree, []byte, map[string]*stack, error) {
	if content.IsGrammar() {
		external := content.Grammar()
		tree, err := app.external(external, isReverse, prevData, currentData)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		if tree == nil {
			return nil, nil, nil, nil, nil
		}

		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		return nil, tree, remaining, stackMap, nil
	}

	if content.IsInstance() {
		instance := content.Instance()
		tree, retStack, err := app.instance(instance, stackMap, escape, channels, isReverse, prevData, currentData)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		return nil, tree, remaining, retStack, nil
	}

	if content.IsRecursive() {
		recursive := content.Recursive()
		if stack, ok := stackMap[recursive]; ok {
			tree, retStack, err := app.token(stack.token, stackMap, escape, channels, isReverse, prevData, currentData)
			if err != nil {
				return nil, nil, nil, nil, err
			}

			remaining := []byte{}
			if tree.HasRemaining() {
				remaining = tree.Remaining()
			}

			return nil, tree, remaining, retStack, nil
		}

		str := fmt.Sprintf("the token (name: %s) was expected to be recursive, but it is not in the current stack", recursive)
		return nil, nil, nil, nil, errors.New(str)
	}

	if len(currentData) < 1 {
		return nil, nil, nil, nil, errors.New("there must be at least 1 value in the given data in order to have an element match, 0 provided")
	}

	grValue := content.Value()
	value, remaining, retStack, err := app.elementValue(grValue, stackMap, escape, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return value, nil, remaining, retStack, nil
}

func (app *application) elementValue(value []byte, stackMap map[string]*stack, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Value, []byte, map[string]*stack, error) {
	remaining := currentData
	builder := app.treeValueBuilder.Create()
	if channels != nil {
		prefix, rem, err := app.channels(channels, prevData, remaining)
		if err == nil {
			builder.WithPrefix(prefix)
			remaining = rem
		}
	}

	if len(remaining) < 1 {
		return nil, nil, nil, errors.New("there must be at least 1 value in the given data in order to have an element match, 0 provided")
	}

	if bytes.HasPrefix(remaining, value) {
		ins, err := builder.WithContent(value).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return ins, remaining[1:], stackMap, nil
	}

	return nil, nil, nil, nil
}

func (app *application) instance(instance grammars.Instance, stackMap map[string]*stack, escape grammars.Token, channels []grammars.Channel, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, map[string]*stack, error) {
	if instance.IsToken() {
		token := instance.Token()
		return app.token(token, stackMap, escape, channels, isReverse, prevData, currentData)
	}

	everything := instance.Everything()
	return app.everything(everything, stackMap, isReverse, prevData, currentData)
}

func (app *application) everything(everything grammars.Everything, stackMap map[string]*stack, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, map[string]*stack, error) {
	exception := everything.Exception()
	escape := everything.Escape()
	return app.token(exception, stackMap, escape, nil, !isReverse, prevData, currentData)
}

func (app *application) channels(channels []grammars.Channel, prevData []byte, currentData []byte) (trees.Trees, []byte, error) {
	treeList := []trees.Tree{}
	remaining := currentData
	previousData := prevData

	for {
		beginAmount := len(treeList)
		for _, oneChannel := range channels {
			tree, err := app.channel(oneChannel, previousData, remaining)
			if err != nil {
				continue
			}

			if tree == nil {
				continue
			}

			prefixLength := len(tree.Bytes(true))
			rem := remaining[prefixLength:]
			if len(rem) == len(remaining) {
				continue
			}

			treeList = append(treeList, tree)
			previousData = remaining
			remaining = rem
		}

		if beginAmount == len(treeList) {
			break
		}
	}

	trees, err := app.treesBuilder.Create().WithList(treeList).Now()
	if err != nil {
		return nil, nil, err
	}

	return trees, remaining, nil
}

func (app *application) channel(channel grammars.Channel, prevData []byte, currentData []byte) (trees.Tree, error) {
	token := channel.Token()
	tree, _, err := app.token(token, map[string]*stack{}, nil, nil, false, prevData, currentData)
	if err != nil {
		return nil, err
	}

	if channel.HasCondition() {
		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		condition := channel.Condition()
		isAccepted, err := app.channelCondition(condition, prevData, remaining)
		if err != nil {
			return nil, err
		}

		if !isAccepted {
			return nil, nil
		}
	}

	return tree, nil
}

func (app *application) channelCondition(condition grammars.ChannelCondition, prevData []byte, nextData []byte) (bool, error) {
	isPrevMatch := true
	if condition.HasPrevious() {
		prevToken := condition.Previous()
		tree, _, err := app.token(prevToken, map[string]*stack{}, nil, nil, false, []byte{}, prevData)
		if err != nil {
			return false, err
		}

		isPrevMatch = tree != nil
	}

	isNextMatch := true
	if condition.HasNext() {
		nextToken := condition.Next()
		tree, _, err := app.token(nextToken, map[string]*stack{}, nil, nil, false, []byte{}, nextData)
		if err != nil {
			return false, err
		}

		isNextMatch = tree != nil
	}
	return isPrevMatch && isNextMatch, nil
}
