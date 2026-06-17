package priorityqueue

const InitCapacity = 10

type PriorityQueue[T Comparable[T]] struct {
	arr      []T
	size     int
	capacity int
}

func New[T Comparable[T]]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		arr:      make([]T, InitCapacity),
		size:     0,
		capacity: InitCapacity,
	}
}

func (pq *PriorityQueue[T]) Push(item T) {
	if pq.size == pq.capacity {
		pq.capacity = pq.capacity * 2
		newArr := make([]T, pq.capacity)
		copy(newArr, pq.arr)
		pq.arr = newArr
	}
	pq.arr[pq.size] = item
	heapifyUp(pq.arr, pq.size)
	pq.size = pq.size + 1
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.size
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.IsEmpty() {
		return *new(T), false
	}
	return pq.arr[0], true
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if pq.IsEmpty() {
		return *new(T), false
	}
	item := pq.arr[0]
	pq.arr[0] = pq.arr[pq.size-1]
	pq.size = pq.size - 1
	heapifyDown(pq.arr, 0, pq.size-1)
	return item, true
}

func heapifyUp[T Comparable[T]](arr []T, idx int) {
	if idx == 0 {
		return
	}
	parentIdx := getParentIdx(idx)
	if arr[idx].CompareTo(arr[parentIdx]) <= 0 {
		// current node has lower or equal priority than parent node
		return
	}
	// current node has higher priority than parent node - swap
	temp := arr[idx]
	arr[idx] = arr[parentIdx]
	arr[parentIdx] = temp
	heapifyUp(arr, parentIdx)
}

func heapifyDown[T Comparable[T]](arr []T, idx int, size int) {
	if idx >= size {
		return
	}
	leftIdx, rightIdx := getChildIdxes(idx)

	highPriorityIdx := idx

	if arr[highPriorityIdx].CompareTo(arr[leftIdx]) <= 0 {
		highPriorityIdx = leftIdx
	}
	if arr[highPriorityIdx].CompareTo(arr[rightIdx]) <= 0 {
		highPriorityIdx = rightIdx
	}

	if idx == highPriorityIdx {
		return
	}

	temp := arr[idx]
	arr[idx] = arr[highPriorityIdx]
	arr[highPriorityIdx] = temp

	heapifyDown(arr, highPriorityIdx, size)
}

func getParentIdx(idx int) int {
	return (idx - 1) / 2
}

func getChildIdxes(idx int) (int, int) {
	return (idx*2 + 1), (idx*2 + 2)
}
