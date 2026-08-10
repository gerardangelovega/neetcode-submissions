type TreeMapNode struct {
	key int
	val int
	left *TreeMapNode
	right *TreeMapNode
}

type TreeMap struct {
	root *TreeMapNode
}

func NewTreeMap() *TreeMap {
	return &TreeMap{}
}

func (tm *TreeMap) Insert(key, val int) {
	if tm.root == nil {
		tm.root = &TreeMapNode{
			key: key,
			val: val,
		}
		return
	}

	curr := tm.root

	for curr != nil {
		if key < curr.key {
			if curr.left == nil {
				curr.left = &TreeMapNode{
					key: key,
					val: val,
				}
				break
			} 
			curr = curr.left
		} else if key > curr.key {
			if curr.right == nil {
				curr.right = &TreeMapNode{
					key: key,
					val: val,
				}
				break
			}
			curr = curr.right
		} else {
			curr.val = val
			break
		}
	}
}

func (tm *TreeMap) Get(key int) int {
	if tm.root == nil {
		return -1
	}

	curr := tm.root

	for curr != nil {
		if key < curr.key {
			curr = curr.left
		} else if key > curr.key {
			curr = curr.right
		} else {
			return curr.val
		}
	}

	return -1
}

func (tm *TreeMap) GetMin() int {
	if tm.root == nil {
		return -1
	}

	return minNode(tm.root).val
}

func (tm *TreeMap) GetMax() int {
	if tm.root == nil {
		return -1
	}

	return maxNode(tm.root).val
}

func (tm *TreeMap) Remove(key int) {
	if tm.root == nil {
		return
	}

	tm.root = removeNode(tm.root, key)
}

func (tm *TreeMap) GetInorderKeys() []int {
	keys := make([]int, 0)

	inorderTraverse(tm.root, &keys)

	return keys
}

func inorderTraverse(root* TreeMapNode, keys *[]int) {
	if root == nil {
		return
	}

	inorderTraverse(root.left, keys)
	*keys = append(*keys, root.key)
	inorderTraverse(root.right, keys)
}

func removeNode(root *TreeMapNode, key int) *TreeMapNode {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = removeNode(root.left, key)
	} else if key > root.key {
		root.right = removeNode(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			minimum := minNode(root.right)
			root.key = minimum.key
			root.val = minimum.val
			root.right = removeNode(root.right, minimum.key)
		}
	}

	return root
}

func minNode(root *TreeMapNode) *TreeMapNode {
	if root == nil {
		return nil
	}

	curr := root

	for curr != nil && curr.left != nil {
		curr = curr.left
	}	

	return curr
}

func maxNode(root *TreeMapNode) *TreeMapNode {
	if root == nil {
		return nil
	}

	curr := root

	for curr != nil && curr.right != nil {
		curr = curr.right
	}	

	return curr
}