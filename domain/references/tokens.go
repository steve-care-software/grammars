package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type tokens struct {
	list []Token
	mp   map[string]Token
}

func createTokens(
	list []Token,
	mp map[string]Token,
) Tokens {
	out := tokens{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the tokens list
func (obj *tokens) List() []Token {
	return obj.list
}

// Fetch fetches a token by hash
func (obj *tokens) Fetch(hash hash.Hash) (Token, error) {
	hashStr := hash.String()
	if ins, ok := obj.mp[hashStr]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the hash (name: %s) do not reference any token instance", hashStr)
	return nil, errors.New(str)
}
