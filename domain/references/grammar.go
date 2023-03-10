package references

import grammars "github.com/steve-care-software/grammars/domain"

type grammar struct {
	name      string
	reference grammars.Grammar
}

func createGrammar(
	name string,
	reference grammars.Grammar,
) Grammar {
	out := grammar{
		name:      name,
		reference: reference,
	}

	return &out
}

// Name returns the name
func (obj *grammar) Name() string {
	return obj.name
}

// Reference returns the reference
func (obj *grammar) Reference() grammars.Grammar {
	return obj.reference
}
