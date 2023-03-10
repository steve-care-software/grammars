package trees

import "github.com/steve-care-software/libs/cryptography/hash"

type content struct {
	value Value
	tree  Tree
}

func createContentWithValue(
	value Value,
) Content {
	return createContentInternally(value, nil)
}

func createContentWithTree(
	tree Tree,
) Content {
	return createContentInternally(nil, tree)
}

func createContentInternally(
	value Value,
	tree Tree,
) Content {
	out := content{
		value: value,
		tree:  tree,
	}

	return &out
}

// Bytes returns the content's bytes
func (obj *content) Bytes(includeChannels bool) []byte {
	if obj.IsValue() {
		return obj.Value().Content()
	}

	return obj.Tree().Bytes(includeChannels)
}

// Hash returns the hash, if any
func (obj *content) Hash() hash.Hash {
	if obj.IsValue() {
		return obj.Value().Hash()
	}

	return obj.Tree().Hash()
}

// IsValue returns true if there is a value, false otherwise
func (obj *content) IsValue() bool {
	return obj.value != nil
}

// Value returns the value if any
func (obj *content) Value() Value {
	return obj.value
}

// IsTree returns true if there is a tree, false otherwise
func (obj *content) IsTree() bool {
	return obj.tree != nil
}

// Tree returns the tree if any
func (obj *content) Tree() Tree {
	return obj.tree
}
