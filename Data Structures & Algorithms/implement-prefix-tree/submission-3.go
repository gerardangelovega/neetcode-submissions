type PrefixTreeNode struct {
	children [26]*PrefixTreeNode
	word bool
}
func NewPrefixTreeNode() *PrefixTreeNode {
	return &PrefixTreeNode{
		word: false,
	}
}

type PrefixTree struct {
	root *PrefixTreeNode
}
func Constructor() PrefixTree {
    return PrefixTree{
		root: NewPrefixTreeNode(),
	}
}
func (this *PrefixTree) Insert(word string) {
	current := this.root
	for _, c := range word {
		index := byte(c) - byte('a')
		if current.children[index] == nil { current.children[index] = NewPrefixTreeNode() }
		child := current.children[index]
		current = child
	}
	current.word = true
}
func (this *PrefixTree) Search(word string) bool {
	current := this.root
	for _, c := range word {
		index := byte(c) - byte('a')
		child := current.children[index]
		if child == nil { return false }
		current = child
	}
	return current.word
}
func (this *PrefixTree) StartsWith(prefix string) bool {
	current := this.root
	for _, c := range prefix {
		index := byte(c) - byte('a')
		child := current.children[index]
		if child == nil { return false }
		current = child
	}
	return true
}
