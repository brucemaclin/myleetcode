package main

type Trie struct {
	root *node
}
type node struct {
	term     bool
	children map[rune]*node
	parent   *node
	val      rune
	meta     interface{}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	n := &node{children: make(map[rune]*node)}
	return Trie{root: n}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	runes := []rune(word)
	var n = this.root
	for _, r := range runes {
		get, ok := n.children[r]
		if ok {
			n = get
		} else {
			newNode := &node{children: make(map[rune]*node)}
			n.children[r] = newNode
			newNode.val = r
			newNode.parent = n
			n = newNode
		}
	}
	newNode := &node{children: make(map[rune]*node)}
	newNode.term = true
	newNode.parent = n
	n.children[0x0] = newNode
	n = newNode
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	runes := []rune(word)
	n := this.root
	for _, v := range runes {
		get, ok := n.children[v]
		if !ok {
			return false
		}
		n = get
	}
	node, ok := n.children[0x0]
	return ok && node.term
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	runes := []rune(prefix)
	n := this.root
	for _, v := range runes {
		get, ok := n.children[v]
		if !ok {
			return false
		}
		n = get
	}
	return n != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
