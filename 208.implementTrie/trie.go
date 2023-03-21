package trie

type Trie struct {
	children map[rune]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	next := t
	for _, char := range word {
		if _, ok := next.children[char]; !ok {
			if next.children == nil {
				next.children = make(map[rune]*Trie)
			}
			next.children[char] = &Trie{}
		}
		child := next.children[char]
		next = child
	}
	next.isWord = true
}

func (t *Trie) Search(word string) bool {
	next := t
	for _, char := range word {
		child, ok := next.children[char]
		if !ok {
			return false
		}
		next = child
	}
	return next.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	next := t
	for _, char := range prefix {
		child, ok := next.children[char]
		if !ok {
			return false
		}
		next = child
	}
	return true
}
