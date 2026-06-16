package priorityqueue

type Comparable[T any] interface {
	CompareTo(T) int
}
