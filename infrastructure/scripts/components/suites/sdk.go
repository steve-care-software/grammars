package suites

import grammars "github.com/steve-care-software/grammars/domain"

// NewSuite creates the suite instance
func NewSuite() Suite {
	suiteBuilder := grammars.NewSuiteBuilder()
	return createSuite(
		suiteBuilder,
	)
}

// Suite represents the suites instance
type Suite interface {
	Suites(values map[string]bool) []grammars.Suite
}
