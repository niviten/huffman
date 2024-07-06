package priorityqueue

type heap[T any] struct {
	arr      []T
	size     int
	capacity int
	compare  Compare[T]
}

func newHeap[T any](compare Compare[T]) *heap[T] {
	return &heap[T]{
		arr:      make([]T, 10),
		size:     0,
		capacity: 10,
		compare:  compare,
	}
}

func (h *heap[T]) insert(item T) {
	h.arr[h.size] = item
	idx := h.size
	h.size = h.size + 1
	if h.size == h.capacity {
		h.capacity = h.capacity * 2
		newArr := make([]T, h.capacity)
		copy(newArr, h.arr)
		h.arr = newArr
	}

	for idx > 0 {
		parentIdx := parentIndex(idx)
		c := h.compare(h.arr[idx], h.arr[parentIdx])
		if c > 0 {
			h.arr[idx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *heap[T]) remove() (T, bool) {
	if h.isEmpty() {
		return *new(T), false
	}
	item := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size = h.size - 1
	h.heapify(0)
	return item, true
}

func (h *heap[T]) peek() (T, bool) {
	if h.isEmpty() {
		return *new(T), false
	}
	return h.arr[0], true
}

func (h *heap[T]) isEmpty() bool {
	return h.size == 0
}

func (h *heap[T]) heapify(index int) {
	left := leftChildIndex(index)
	right := rightChildIndex(index)
	largest := index
	if left < h.size && h.compare(h.arr[left], h.arr[largest]) > 0 {
		largest = left
	}
	if right < h.size && h.compare(h.arr[right], h.arr[largest]) > 0 {
		largest = right
	}
	if largest != index {
		h.arr[index], h.arr[largest] = h.arr[largest], h.arr[index]
		h.heapify(largest)
	}
}

func parentIndex(index int) int {
	return (index - 1) / 2
}

func leftChildIndex(index int) int {
	return rightChildIndex(index) - 1
}

func rightChildIndex(index int) int {
	return (index + 1) * 2
}
