package suites

import grammars "github.com/steve-care-software/grammars/domain"

type suite struct {
	suiteBuilder grammars.SuiteBuilder
}

func createSuite(
	suiteBuilder grammars.SuiteBuilder,
) Suite {
	out := suite{
		suiteBuilder: suiteBuilder,
	}

	return &out
}

// Suites creates suites based on the values
func (app *suite) Suites(values map[string]bool) []grammars.Suite {
	list := []grammars.Suite{}
	for str, isValid := range values {
		suite := app.suite([]byte(str), isValid)
		list = append(list, suite)
	}

	if len(list) <= 0 {
		return nil
	}

	return list
}

func (app *suite) suite(values []byte, isValid bool) grammars.Suite {
	builder := app.suiteBuilder.Create()
	if isValid {
		builder.WithValid(values)
	}

	if !isValid {
		builder.WithInvalid(values)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
