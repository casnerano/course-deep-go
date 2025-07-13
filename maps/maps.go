package maps

import "cmp"

type node[K cmp.Ordered, V any] struct {
	key   K
	value V

	left  *node[K, V]
	right *node[K, V]
}

type OrderedMap[K cmp.Ordered, V any] struct {
	Root *node[K, V]
	size int
}

func NewOrderedMap[K cmp.Ordered, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	var insert func(n **node[K, V])
	insert = func(n **node[K, V]) {
		if *n == nil {
			*n = &node[K, V]{key: key, value: value}
			m.size++
			return
		}

		if (*n).key == key {
			(*n).value = value
			return
		}

		if key < (*n).key {
			insert(&(*n).left)
		} else {
			insert(&(*n).right)
		}
	}

	insert(&m.Root)
}

func (m *OrderedMap[K, V]) Erase(key K) {
	findMinNode := func(n *node[K, V]) *node[K, V] {
		for n != nil && n.left != nil {
			n = n.left
		}
		return n
	}

	var erase func(n **node[K, V])
	erase = func(n **node[K, V]) {
		if *n == nil {
			return
		}

		if (*n).key == key {
			if (*n).left == nil {
				*n = (*n).right
			} else if (*n).right == nil {
				*n = (*n).left
			} else {
				minNode := findMinNode((*n).right)
				*n = minNode
			}

			m.size--
			return
		}

		if key < (*n).key {
			erase(&(*n).left)
		} else {
			erase(&(*n).right)
		}
	}

	erase(&m.Root)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	var contains func(n *node[K, V]) bool
	contains = func(n *node[K, V]) bool {
		if n == nil {
			return false
		}

		if n.key == key {
			return true
		} else if key < n.key {
			return contains(n.left)
		} else {
			return contains(n.right)
		}
	}

	return contains(m.Root)
}

func (m *OrderedMap[K, V]) Size() int {
	return m.size
}

func (m *OrderedMap[K, V]) ForEach(action func(key K, value V)) {
	var walk func(n *node[K, V])
	walk = func(n *node[K, V]) {
		if n == nil {
			return
		}

		if n.left != nil {
			walk(n.left)
		}

		action(n.key, n.value)

		if n.right != nil {
			walk(n.right)
		}
	}

	walk(m.Root)
}
