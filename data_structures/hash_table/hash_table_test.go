package main

import (
	"reflect"
	"testing"
)

func TestPut(t *testing.T) {
	h := &HashTable{
		data: make(map[int]string),
	}

	cases := []struct {
		key      int
		value    string
		expected *HashTable
	}{
		{
			key:   1,
			value: "foo",
			expected: &HashTable{
				data: map[int]string{
					1: "foo",
				},
			},
		},
		{
			key:   2,
			value: "bar",
			expected: &HashTable{
				data: map[int]string{
					1: "foo",
					2: "bar",
				},
			},
		},
		{
			key:   3,
			value: "foobar",
			expected: &HashTable{
				data: map[int]string{
					1: "foo",
					2: "bar",
					3: "foobar",
				},
			},
		},
	}

	for _, c := range cases {
		h.put(c.key, c.value)
		if !reflect.DeepEqual(h, c.expected) {
			t.Errorf("key:%v value:%v expected:%v", c.key, c.value, c.expected)
		}
	}
}

func TestGet(t *testing.T) {
	h := &HashTable{
		data: make(map[int]string),
	}

	cases := []struct {
		key      int
		value    string
		expected string
	}{
		{
			key:      1,
			value:    "foo",
			expected: "foo",
		},
		{
			key:      2,
			value:    "bar",
			expected: "bar",
		},
		{
			key:      3,
			value:    "foobar",
			expected: "foobar",
		},
	}

	for _, c := range cases {
		h.put(c.key, c.value)
		if h.get(c.key) != c.expected {
			t.Errorf("key:%v value:%v expected:%v", c.key, c.value, c.expected)
		}
	}
}
