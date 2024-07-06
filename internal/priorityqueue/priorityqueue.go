package priorityqueue

type PriorityQueue[T any] struct {
    h *heap[T]
}

func New [T any] (compare Compare[T]) *PriorityQueue[T] {
    return &PriorityQueue[T]{
        h: newHeap(compare),
    }
}

func (pq *PriorityQueue[T]) Add(item T) {
    pq.h.insert(item)
}

func (pq *PriorityQueue[T]) Take() (T, bool) {
    return pq.h.remove()
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
    return pq.h.peek()
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
    return pq.h.isEmpty()
}
