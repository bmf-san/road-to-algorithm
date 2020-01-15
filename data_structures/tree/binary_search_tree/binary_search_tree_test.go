package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &BST{}

	tree.insert(10)
	tree.insert(2)
	tree.insert(3)
	tree.insert(15)

	expected := &BST{
		Root: &Node{
			Key: 10,
			Left: &Node{
				Key:  2,
				Left: nil,
				Right: &Node{
					Key:   3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Key:   15,
				Left:  nil,
				Right: nil,
			},
		},
	}

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("actual: %v expected: %v", tree, expected)
	}
}

func TestRemove(t *testing.T) {
	tree := &BST{}

	tree.insert(10)
	tree.insert(2)
	tree.insert(3)
	tree.insert(15)

	tree.remove(3)

	expected := &BST{
		Root: &Node{
			Key: 10,
			Left: &Node{
				Key:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Key:   15,
				Left:  nil,
				Right: nil,
			},
		},
	}

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("actual: %v expected: %v", tree, expected)
	}
}

func TestSearch(t *testing.T) {
	tree := &BST{}

	tree.insert(10)
	tree.insert(2)
	tree.insert(3)
	tree.insert(3)
	tree.insert(3)
	tree.insert(15)
	tree.insert(14)
	tree.insert(18)
	tree.insert(16)
	tree.insert(16)

	cases := []struct {
		item     int
		expected bool
	}{
		{
			item:     10,
			expected: true,
		},
		{
			item:     19,
			expected: false,
		},
	}

	for _, c := range cases {
		actual := tree.search(c.item)
		if actual != c.expected {
			t.Errorf("actual: %v expected: %v", actual, c.expected)
		}
	}
}
