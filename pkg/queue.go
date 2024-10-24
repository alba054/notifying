package pkg

// this is a classic queue data structure implementation
type Queue[T interface{}] struct {
	data []T
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.data) < 1 {
		return nil
	}

	value := &q.data[0]

	if len(q.data) == 1 {
		q.data = []T{}
	} else {
		q.data = q.data[1:]
	}

	return value
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}
