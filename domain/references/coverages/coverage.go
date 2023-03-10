package coverages

import "github.com/steve-care-software/grammars/domain/references"

type coverage struct {
	token      references.Token
	executions Executions
}

func createCoverage(
	token references.Token,
	executions Executions,
) Coverage {
	out := coverage{
		token:      token,
		executions: executions,
	}

	return &out
}

// Token returns the token
func (obj *coverage) Token() references.Token {
	return obj.token
}

// Executions returns the executions
func (obj *coverage) Executions() Executions {
	return obj.executions
}
