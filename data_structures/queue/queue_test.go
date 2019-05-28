package main

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	actual := newQueue()
	expected := &Queue{}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestEnqueue(t *testing.T) {
	actual := newQueue()
	actual.enqueue(&Node{
		value: "first",
	})
	actual.enqueue(&Node{
		value: "second",
	})
	actual.enqueue(&Node{
		value: "third",
	})

	expected := &Queue{
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

func TestDequeue(t *testing.T) {
	actual := newQueue()
	actual.enqueue(&Node{
		value: "first",
	})
	actual.enqueue(&Node{
		value: "second",
	})
	actual.enqueue(&Node{
		value: "third",
	})
	actual.dequeue()

	expected := &Queue{
		nodes: []*Node{
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
