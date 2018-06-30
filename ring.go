package consistenthashing

// key value is nodeName, hashed using 32 bit fnv1a
// consistenthashing use bst data structure

import (
	"errors"
)

var (
	errorsNoNodeAssigned = errors.New("cannot found assigned node")
)

// Ring yield bst ring
type Ring struct {
	bst *bst
}

// Init initialise ring
func Init() *Ring {
	return &Ring{
		bst: &bst{},
	}
}

// Insert inserting node and node data to ring
func (r *Ring) Insert(nodeName string, nodeData interface{}) {

	r.bst.insert(hash(nodeName), nodeData)
}

// Remove removing element based on node name
func (r *Ring) Remove(nodeName string) {
	r.bst.remove(hash(nodeName))
}

// Get get assigned node
func (r *Ring) Get(key string) (interface{}, error) {
	hashed := hash(key)
	a := r.bst.findNextGreater(hashed)

	if a == nil {
		a = r.bst.findPrevSmaller(hashed)
		if a == nil {
			return nil, errorsNoNodeAssigned
		}
	}
	return a.val, nil
}

// GetAllElmt get element of inserted val
func (r *Ring) GetAllElmt() []interface{} {
	var res []interface{}
	nodes := r.bst.inorder()
	for _, v := range nodes {
		res = append(res, v.val)
	}
	return res
}
