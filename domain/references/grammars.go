package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type grammarsStr struct {
	list []Grammar
	mp   map[string]Grammar
}

func createGrammars(
	list []Grammar,
	mp map[string]Grammar,
) Grammars {
	out := grammarsStr{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the grammars list
func (obj *grammarsStr) List() []Grammar {
	return obj.list
}

// Fetch fetches a grammar by hash
func (obj *grammarsStr) Fetch(hash hash.Hash) (Grammar, error) {
	hashStr := hash.String()
	if ins, ok := obj.mp[hashStr]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the hash (name: %s) do not reference any grammar instance", hashStr)
	return nil, errors.New(str)
}
