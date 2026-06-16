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

func (pq *PriorityQueue[T]) Insert(item T) {
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

func getParentIdx(idx int) int {
	return (idx - 1) / 2
}
