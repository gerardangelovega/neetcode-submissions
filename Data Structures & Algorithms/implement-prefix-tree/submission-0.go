type PrefixTreeNode struct {
    children map[rune]*PrefixTreeNode
    word bool
}

func NewPrefixTreeNode() *PrefixTreeNode {
    return &PrefixTreeNode {
        children: make(map[rune]*PrefixTreeNode),
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
    curr := this.root
    for _, r := range word {
        if _, e := curr.children[r]; !e {
            curr.children[r] = NewPrefixTreeNode()
        }
        curr = curr.children[r]
    }
    curr.word = true
}

func (this *PrefixTree) Search(word string) bool {
    curr := this.root
    for _, r := range word {
        if _, e := curr.children[r]; !e {
            return false;
        }
        curr = curr.children[r]
    }
    return curr.word
}

func (this *PrefixTree) StartsWith(prefix string) bool {
    curr := this.root
    for _, r := range prefix {
        if _, e := curr.children[r]; !e {
            return false;
        }
        curr = curr.children[r]
    }
    return true
}
