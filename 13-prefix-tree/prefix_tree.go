package p13prefixtree

type Node struct {
	next  [128]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{&Node{}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.next[c] == nil {
			node.next[c] = &Node{}
		}
		node = node.next[c]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	if t.root == nil {
		return false
	}
	node := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		node = node.next[c]
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	if t.root == nil {
		return false
	}
	node := t.root
	for i := 0; i < len(prefix); i++ {
		if node == nil {
			return false
		}
		c := prefix[i]
		if node.next[c] == nil {
			return false
		}
		node = node.next[c]
	}
	return true
}
