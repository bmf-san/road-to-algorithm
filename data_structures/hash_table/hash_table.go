package main

import "fmt"

// A HashTable is hash table.
type HashTable struct {
	data map[int]string
}

// hash is create a hash key.
func hash(key int) int {
	return key % 5
}

// put is add key to hash table.
func (h HashTable) put(key int, value string) {
	hash := hash(key)
	if h.data == nil {
		h.data = make(map[int]string)
	}
	h.data[hash] = value
}

// get is get a value from hash table.
func (h HashTable) get(key int) string {
	var hash int = hash(key)

	return h.data[hash]
}

func main() {
	h := &HashTable{
		data: make(map[int]string),
	}

	h.put(1, "foo")
	h.put(2, "bar")

	fmt.Printf("%#v\n", h.get(1))
	fmt.Printf("%#v\n", h.get(2))
}
