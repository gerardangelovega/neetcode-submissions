type Queue struct {
	items []*WordDictionaryNode
}

func NewQueue() Queue {
	return Queue{
		items: make([]*WordDictionaryNode, 0),
	}
}

func (q *Queue) Enqueue(val *WordDictionaryNode) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() *WordDictionaryNode {
	res := q.items[0]
	q.items = q.items[1:]
	return res
}

func (q *Queue) Length() int {
	return len(q.items)
}

type WordDictionaryNode struct {
	children map[rune]*WordDictionaryNode
	word bool
}

func NewWordDictionaryNode() *WordDictionaryNode {
	return &WordDictionaryNode{
		children: make(map[rune]*WordDictionaryNode),
		word: false,
	}
}

type WordDictionary struct {
	root *WordDictionaryNode 
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewWordDictionaryNode(),
	}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.root
	for _, r := range word {
		if _, e := curr.children[r]; !e {
			curr.children[r] = NewWordDictionaryNode()
		}
		curr = curr.children[r]
	}
	curr.word = true
}

func (this *WordDictionary) Search(word string) bool {
	return search(this.root, word)
}

func search(root *WordDictionaryNode, word string) bool {
	r := rune(word[0])
	q := NewQueue()
	if r != '.' {
		if _, e := root.children[r]; !e {
			return false
		}
		if len(word) == 1 {
			return root.children[r].word
		}
		q.Enqueue(root.children[r])
	} else {
		if len(word) == 1 {
			for _, v := range root.children {
				if v.word {
					return true
				}
			}
			return false
		}
		for _, v := range root.children {
			q.Enqueue(v)
		}
	}

	for q.Length() != 0 {
		curr := q.Dequeue()
		if search(curr, word[1:]) {
			return true
		}
	}
	
	return false
}