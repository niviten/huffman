package priorityqueue

type Compare[T any] func(a, b T) int
