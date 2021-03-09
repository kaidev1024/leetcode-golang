type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	isWord   bool
	children []*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isWord:   false,
		children: make([]*TrieNode, 26),
	}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{NewTrieNode()}
}

func (this *WordDictionary) AddWord(word string) {
	root := this.root
	for _, char := range word {
		index := char - 'a'
		if root.children[index] == nil {
			root.children[index] = NewTrieNode()
		}
		root = root.children[index]
	}
	root.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	root := this.root
	var helper func(root *TrieNode, word string) bool
	helper = func(root *TrieNode, word string) bool {
		if root == nil {
			return false
		}
		if word == "" {
			return root.isWord
		}
		char := word[0]
		if char == '.' {
			for i := 0; i < 26; i++ {
				if root.children[i] != nil && helper(root.children[i], word[1:]) {
					return true
				}
			}
			return false
		}
		index := char - 'a'
		if root.children[index] == nil {
			return false
		}
		return helper(root.children[index], word[1:])
	}
	return helper(root, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */