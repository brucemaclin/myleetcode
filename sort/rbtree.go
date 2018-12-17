package main

type Node struct {
	Left  *Node
	Right *Node
	Item
	Black bool
}
type LLRB struct {
	count int
	root  *Node
}
type Item interface {
	Less(than Item) bool
}

func New() *LLRB {
	return &LLRB{}
}

func (l *LLRB) SetRoot(root *Node) {
	l.root = root
}

func (l *LLRB) Root() *Node {
	return l.root
}
func (l *LLRB) Len() int {
	return l.count
}
func (l *LLRB) InsertOrReplace(item Item) Item {
	if item == nil {
		panic("item should not be nil")
	}
	var replaced Item
	l.root, replaced = l.insertOrReplace(l.root, item)
	if replaced == nil {
		l.count++
	}
	l.root.Black = true
	return replaced
}
func less(x Item, y Item) bool {
	return x.Less(y)
}
func (l *LLRB) insertOrReplace(h *Node, item Item) (*Node, Item) {
	if h == nil {
		return newNode(item), nil
	}
	var replaced Item
	if less(h.Item, item) {
		h.Right, replaced = l.insertOrReplace(h.Right, item)
	} else if less(item, h.Item) {
		h.Left, replaced = l.insertOrReplace(h.Left, item)
	} else {
		replaced, h.Item = h.Item, item
	}
	if !isRed(h.Left) && isRed(h.Right) {
		h = rotateLeft(h)
	}
	if isRed(h.Left) && isRed(h.Left.Left) {
		h = rotateRight(h)
	}
	if isRed(h.Left) && isRed(h.Right) {
		flip(h)
	}
	return h, replaced
}
func newNode(item Item) *Node {
	return &Node{Item: item}
}

func rotateLeft(h *Node) *Node {
	x := h.Right
	x.Black = h.Black
	h.Right = x.Left
	x.Left = h
	h.Black = false
	return x
}
func rotateRight(h *Node) *Node {
	x := h.Left
	h.Left = x.Right

	x.Black = h.Black
	x.Right = h
	h.Black = false
	return x

}
func flip(h *Node) {
	h.Black = !h.Black
	h.Left.Black = !h.Left.Black
	h.Right.Black = !h.Right.Black
}
func isRed(node *Node) bool {
	return !node.Black
}
