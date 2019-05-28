package main

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	actual := newStack()
	expected := &Stack{}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestPush(t *testing.T) {
	actual := newStack()
	actual.push(&Node{
		value: "first",
	})
	actual.push(&Node{
		value: "second",
	})
	actual.push(&Node{
		value: "third",
	})

	expected := &Stack{
		nodes: []*Node{
			&Node{
				value: "first",
			},
			&Node{
				value: "second",
			},
			&Node{
				value: "third",
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestPop(t *testing.T) {
	actual := newStack()
	actual.push(&Node{
		value: "first",
	})
	actual.push(&Node{
		value: "second",
	})
	actual.push(&Node{
		value: "third",
	})
	actual.pop()

	expected := &Stack{
		nodes: []*Node{
			&Node{
				value: "first",
			},
			&Node{
				value: "second",
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}
