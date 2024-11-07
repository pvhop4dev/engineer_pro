package leetcode

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isEnd    bool
	children [26]*TrieNode
}

func ConstructorTrie() Trie {
	return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = &TrieNode{}
		}
		cur = cur.children[c-'a']
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return cur.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, c := range prefix {
		if cur.children[c-'a'] == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return true
}
