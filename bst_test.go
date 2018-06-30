package consistenthashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBst(t *testing.T) {
	assert := assert.New(t)
	bst := &bst{}

	datas := []int{5, 7, 3, 6, 12, 20, 1, 2, 9, 10}

	for _, v := range datas {
		bst.insert(uint32(v), nil)
	}

	preorderNode := bst.preorder()
	assert.NotNil(preorderNode, "preordernode should not be nil")
	assert.Equal(len(datas), len(preorderNode))

	inorderNode := bst.inorder()
	assert.NotNil(inorderNode, "inordernode should not be nil")
	assert.Equal(len(datas), len(inorderNode))

	postorderNode := bst.postorder()
	assert.NotNil(postorderNode, "postordernode should not be nil")
	assert.Equal(len(datas), len(postorderNode))

	// test predecessor and successor
	node12 := bst.find(uint32(12))
	assert.NotNil(node12)
	assert.Equal(uint32(12), node12.key)

	prevNode12 := bst.predecessor(node12)
	assert.NotNil(prevNode12)
	assert.Equal(uint32(10), prevNode12.key)

	nextNode12 := bst.successor(node12)
	assert.NotNil(nextNode12)
	assert.Equal(uint32(20), nextNode12.key)

	n := bst.findNextGreater(uint32(13))
	assert.NotNil(n)
	assert.Equal(uint32(20), n.key)

	nprv := bst.findPrevSmaller(uint32(13))
	assert.NotNil(nprv)
	assert.Equal(uint32(12), nprv.key)

	bst.remove(uint32(3))

	result := bst.find(uint32(3))
	assert.Nil(result)

	inorderNode = bst.inorder()
	assert.NotNil(inorderNode, "inordernode should not be nil")
	assert.Equal(len(datas)-1, len(inorderNode))
}

func TestDeleteNode(t *testing.T) {
	assert := assert.New(t)
	bst := &bst{}

	datas := []int{6, 7, 3, 2, 4}

	for _, v := range datas {
		bst.insert(uint32(v), nil)
	}

	// delete case does not have a child
	bst.remove(uint32(2))
	result := bst.find(uint32(2))
	assert.Nil(result)

	// delete case have one child
	bst.remove(uint32(4))
	result = bst.find(uint32(4))
	assert.Nil(result)

	// delete root / have both child
	bst.remove(uint32(6))
	result = bst.find(uint32(6))
	assert.Nil(result)

	// now 7 become root since successor
	bst.remove(uint32(7))
	result = bst.find(uint32(7))
	assert.Nil(result)
}
