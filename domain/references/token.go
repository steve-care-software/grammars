package references

import grammars "github.com/steve-care-software/grammars/domain"

type token struct {
	name      string
	reference grammars.Token
}

func createToken(
	name string,
	reference grammars.Token,
) Token {
	out := token{
		name:      name,
		reference: reference,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Reference returns the reference
func (obj *token) Reference() grammars.Token {
	return obj.reference
}
