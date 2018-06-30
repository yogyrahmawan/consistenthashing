package consistenthashing

type node struct {
	key    uint32
	val    interface{}
	parent *node
	left   *node
	right  *node
}

type bst struct {
	root *node
}

func (b *bst) insert(key uint32, val interface{}) {

	newNode := &node{key: key, val: val, parent: nil, left: nil, right: nil}
	if b.root == nil {
		b.root = newNode
	} else {
		insertNode(b.root, newNode)
	}
}

func insertNode(root, newNode *node) {
	if newNode.key <= root.key {
		if root.left != nil {
			insertNode(root.left, newNode)
		} else {
			root.left = newNode
			root.left.parent = root
		}
	} else {
		if root.right != nil {
			insertNode(root.right, newNode)
		} else {
			root.right = newNode
			root.right.parent = root
		}
	}
}

func (b *bst) find(key uint32) *node {
	return findNode(b.root, key)
}

func findNode(root *node, key uint32) *node {
	if root == nil {
		return nil
	}

	if key == root.key {
		return root
	}

	if key < root.key {
		return findNode(root.left, key)
	}

	if key > root.key {
		return findNode(root.right, key)
	}

	return root
}

func (b *bst) min(node *node) *node {
	lnode := node
	if lnode == nil {
		return nil
	}

	for lnode.left != nil {
		lnode = lnode.left
	}

	return lnode
}

func (b *bst) max(node *node) *node {
	lnode := node
	if lnode == nil {
		return nil
	}

	for lnode.right != nil {
		lnode = lnode.right
	}
	return lnode
}

func (b *bst) remove(key uint32) {
	b.removeNode(key)
}

func (b *bst) findPrevSmaller(key uint32) *node {

	node := b.find(key)

	if node == nil {
		return b.prevClosest(b.root, key)
	}
	smallerNode := b.predecessor(node)
	if smallerNode != nil {
		return smallerNode
	}

	return node
}

func (b *bst) prevClosest(node *node, key uint32) *node {
	if node == nil {
		return nil
	}

	prevNode := node
	for node != nil {
		prevNode = node
		if node.key < key {
			node = node.right
		} else if node.key > key {
			node = node.left
		} else {
			return b.predecessor(node)
		}
	}

	if key > prevNode.key {
		return prevNode
	}

	return b.predecessor(prevNode)
}

func (b *bst) findNextGreater(key uint32) *node {

	node := b.find(key)
	if node == nil {
		return b.nextClosesElmt(b.root, key)
	}
	sNode := b.successor(node)
	if node != nil {
		return sNode
	}
	return node
}

func (b *bst) nextClosesElmt(root *node, key uint32) *node {
	if root == nil {
		return nil
	}

	prevNode := root
	for root != nil {
		prevNode = root
		if root.key < key {
			root = root.right
		} else if root.key > key {
			root = root.left
		} else {
			return b.successor(root)
		}
	}

	if key < prevNode.key {
		return prevNode
	}

	return b.successor(prevNode)
}

func (b *bst) removeNode(key uint32) bool {
	node := b.find(key)
	if node == nil {
		return false
	}

	// node does not have child
	if node.left == nil && node.right == nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = nil
			} else if node.parent.right == node {
				node.parent.right = nil
			}
		} else {
			b.root = nil
		}
	} else if node.left == nil || node.right == nil {
		// node have only one child
		childNode := node.right
		if node.right == nil {
			childNode = node.left
		}

		parent := node.parent
		if parent != nil {
			if parent.left == node {
				parent.left = childNode
			} else if parent.right == node {
				parent.right = childNode
			}
		} else {
			childNode.parent = parent
			b.root = childNode
		}
	} else {
		// have both child
		successor := b.successor(node)

		// remove successor from it's parent
		successorParent := successor.parent
		if successorParent.left == successor {
			successorParent.left = nil
		} else {
			successorParent.right = nil
		}

		node.key = successor.key
		node.val = successor.val
	}

	return true
}

// the predecessor of node x is node with greatest key smaller than x
func (b *bst) predecessor(node *node) *node {
	if node.left != nil {
		return b.max(node.left)
	}

	currentNode := node
	parent := currentNode.parent
	if parent != nil && parent.left == currentNode {
		currentNode = parent
		parent = currentNode.parent
	}

	return parent
}

// the successor of node x is node with minimum key greater than x
func (b *bst) successor(node *node) *node {
	if node.right != nil {
		return b.min(node.right)
	}

	if node.parent == nil {
		return node
	}

	currentNode := node
	parent := currentNode.parent
	for parent != nil && currentNode == parent.right {
		currentNode = parent
		parent = currentNode.parent
	}

	return parent

}

// left root right
func (b *bst) inorder() []*node {

	var res []*node
	return iterateInorder(b.root, res)
}

func iterateInorder(root *node, res []*node) []*node {
	if root == nil {
		return res
	}

	if root.left != nil {
		res = iterateInorder(root.left, res)
	}

	res = append(res, root)

	if root.right != nil {
		res = iterateInorder(root.right, res)
	}

	return res
}

// root, left , right
func (b *bst) preorder() []*node {

	var res []*node
	return iteratePreorder(b.root, res)
}

func iteratePreorder(root *node, res []*node) []*node {
	if root == nil {
		return res
	}

	res = append(res, root)

	if root.left != nil {
		res = iteratePreorder(root.left, res)

	}

	if root.right != nil {
		res = iteratePreorder(root.right, res)

	}

	return res
}

func (b *bst) postorder() []*node {

	var res []*node
	return iteratePostorder(b.root, res)
}

// left, right , root
func iteratePostorder(root *node, res []*node) []*node {
	if root == nil {
		return res
	}

	if root.left != nil {
		res = iteratePostorder(root.left, res)
	}

	if root.right != nil {
		res = iteratePostorder(root.right, res)
	}

	return append(res, root)
}
