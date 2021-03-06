package node

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/begmaroman/mpt/enc"
)

func TestBranchNode_Find(t *testing.T) {
	key, value := enc.BytesToHex([]byte("key")), []byte("value")

	nd := NewBranchNode()
	nd.Children[key[0]] = NewExtensionNode(key[1:], NewLeafNode(value))

	rVal, rNd, ok := nd.Find(key)
	if !ok {
		t.Error("added data not found")
		return
	}

	if !bytes.Equal(rVal, value) {
		t.Errorf("invalid value: expected %s got %s", value, rVal)
	}

	if !reflect.DeepEqual(nd, rNd) {
		t.Errorf("invalid node: expected %#v got %#v", nd, rNd)
	}
}

func TestBranchNode_PutAndCheck(t *testing.T) {
	key, value := enc.BytesToHex([]byte("key")), []byte("value")

	nd := NewBranchNode()
	newNode, ok := nd.Put(key, NewLeafNode(value))
	if !ok {
		t.Error("put failed")
		return
	}

	rVal, rNd, ok := newNode.Find(key)
	if !ok {
		t.Error("added data not found")
		return
	}

	if !bytes.Equal(rVal, value) {
		t.Errorf("invalid value: expected %s got %s", value, rVal)
	}

	if !reflect.DeepEqual(newNode, rNd) {
		t.Errorf("invalid node: expected %#v got %#v", newNode, rNd)
	}
}

func TestBranchNode_Cache(t *testing.T) {
	nd := NewBranchNode()
	nd.Hash = []byte("test")
	nd.Dirty = true

	hash, dirty := nd.Cache()
	if !bytes.Equal(hash, nd.Hash) {
		t.Errorf("invalid hash: expected %s got %s", nd.Hash, hash)
	}

	if nd.Dirty != dirty {
		t.Errorf("invalid dirty: expected %t got %t", nd.Dirty, dirty)
	}
}
