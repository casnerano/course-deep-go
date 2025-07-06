package slices_and_arrays

type CircularQueue[T any] struct {
	values []T
	front  int
	rear   int
	size   int
}

func NewCircularQueue[T any](qSize int) *CircularQueue[T] {
	return &CircularQueue[T]{
		values: make([]T, qSize),
		front:  -1,
		rear:   -1,
		size:   0,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	if q.front == -1 {
		q.front = 0
	}

	q.rear = (q.rear + 1) % len(q.values)
	q.values[q.rear] = value
	q.size++

	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	q.front = (q.front + 1) % len(q.values)
	q.size--

	return true
}

func (q *CircularQueue[T]) Front() (T, bool) {
	if q.Empty() {
		var zero T
		return zero, false
	}

	return q.values[q.front], true
}

func (q *CircularQueue[T]) Back() (T, bool) {
	if q.Empty() {
		var zero T
		return zero, false
	}

	return q.values[q.rear], true
}

func (q *CircularQueue[T]) Empty() bool {
	return q.size < 1
}

func (q *CircularQueue[T]) Full() bool {
	return q.size == len(q.values)
}
