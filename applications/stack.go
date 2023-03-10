package applications

import grammars "github.com/steve-care-software/grammars/domain"

type stack struct {
	token grammars.Token
	lines map[int][]byte
}
