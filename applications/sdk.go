package applications

import (
	grammars "github.com/steve-care-software/grammars/domain"
	"github.com/steve-care-software/grammars/domain/references"
	"github.com/steve-care-software/grammars/domain/references/coverages"
	"github.com/steve-care-software/grammars/domain/trees"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	grammarTokenBuilder := grammars.NewTokenBuilder()
	treesBuilder := trees.NewBuilder()
	treeBuilder := trees.NewTreeBuilder()
	treeTokenBuilder := trees.NewTokenBuilder()
	treeLineBuilder := trees.NewLineBuilder()
	treeElementBuilder := trees.NewElementBuilder()
	treeContentBuilder := trees.NewContentBuilder()
	treeValueBuilder := trees.NewValueBuilder()
	coveragesBuilder := coverages.NewBuilder()
	coverageBuilder := coverages.NewCoverageBuilder()
	coverageExecutionsBuilder := coverages.NewExecutionsBuilder()
	coverageExecutionBuilder := coverages.NewExecutionBuilder()
	coverageResultBuilder := coverages.NewResultBuilder()
	return createApplication(
		grammarTokenBuilder,
		treesBuilder,
		treeBuilder,
		treeTokenBuilder,
		treeLineBuilder,
		treeElementBuilder,
		treeContentBuilder,
		treeValueBuilder,
		coveragesBuilder,
		coverageBuilder,
		coverageExecutionsBuilder,
		coverageExecutionBuilder,
		coverageResultBuilder,
	)
}

// Application represents the AST application
type Application interface {
	Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error)
	Coverages(reference references.Reference) (coverages.Coverages, error)
}
