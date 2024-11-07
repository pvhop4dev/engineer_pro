package leetcode

type WordDictionary struct {
	root *WordDictionaryNode
}

type WordDictionaryNode struct {
	isEnd    bool
	children [26]*WordDictionaryNode
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{root: &WordDictionaryNode{}}
}

func (w *WordDictionary) AddWord(word string) {
	cur := w.root
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = &WordDictionaryNode{}
		}
		cur = cur.children[c-'a']
	}
	cur.isEnd = true
}

func (w *WordDictionary) Search(word string) bool {
	return w.search(w.root, word)
}

func (w *WordDictionary) search(node *WordDictionaryNode, word string) bool {
	for i, c := range word {
		if c == '.' {
			for _, child := range node.children {
				if child != nil && w.search(child, word[i+1:]) {
					return true
				}
			}
			return false
		}
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return node.isEnd
}
