package trees

import "github.com/steve-care-software/libs/cryptography/hash"

type trees struct {
	hash hash.Hash
	list []Tree
}

func createTrees(
	hash hash.Hash,
	list []Tree,
) Trees {
	out := trees{
		hash: hash,
		list: list,
	}

	return &out
}

// Bytes returns the trees' bytes
func (obj *trees) Bytes(includeChannels bool) []byte {
	output := []byte{}
	for _, oneTree := range obj.list {
		output = append(output, oneTree.Bytes(includeChannels)...)
	}

	return output
}

// Hash returns the hash, if any
func (obj *trees) Hash() hash.Hash {
	return obj.hash
}

// List returns the trees
func (obj *trees) List() []Tree {
	return obj.list
}
