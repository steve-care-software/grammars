package trees

import "github.com/steve-care-software/libs/cryptography/hash"

type value struct {
	hash    hash.Hash
	content []byte
	prefix  Trees
}

func createValue(
	hash hash.Hash,
	content []byte,
) Value {
	return createValueInternally(hash, content, nil)
}

func createValueWithPrefix(
	hash hash.Hash,
	content []byte,
	prefix Trees,
) Value {
	return createValueInternally(hash, content, prefix)
}

func createValueInternally(
	hash hash.Hash,
	content []byte,
	prefix Trees,
) Value {
	out := value{
		hash:    hash,
		content: content,
		prefix:  prefix,
	}

	return &out
}

// Hash returns the hash, if any
func (obj *value) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *value) Content() []byte {
	return obj.content
}

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *value) HasPrefix() bool {
	return obj.prefix != nil
}

// Prefix returns the prefix, if any
func (obj *value) Prefix() Trees {
	return obj.prefix
}
