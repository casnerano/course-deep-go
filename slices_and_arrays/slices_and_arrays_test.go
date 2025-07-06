package slices_and_arrays

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	var value int
	var ok bool

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	_, ok = queue.Front()
	assert.False(t, ok)

	value, ok = queue.Back()
	assert.False(t, ok)

	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	value, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	value, ok = queue.Back()
	assert.True(t, ok)
	assert.Equal(t, 3, value)

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	value, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	value, ok = queue.Back()
	assert.True(t, ok)
	assert.Equal(t, 4, value)

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
