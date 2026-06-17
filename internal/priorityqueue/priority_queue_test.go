package priorityqueue

import "testing"

type testItem struct {
	priority int
	id       int
}

func (ti testItem) CompareTo(other testItem) int {
	if ti.priority > other.priority {
		return 1
	}
	if ti.priority < other.priority {
		return -1
	}
	return 0
}

func TestPriorityQueuePushAddsFirstItem(t *testing.T) {
	pq := New[testItem]()
	first := testItem{priority: 42, id: 1}

	pq.Push(first)

	if pq.size != 1 {
		t.Fatalf("size = %d, want 1", pq.size)
	}
	if pq.capacity != InitCapacity {
		t.Fatalf("capacity = %d, want %d", pq.capacity, InitCapacity)
	}
	if len(pq.arr) != InitCapacity {
		t.Fatalf("len(arr) = %d, want %d", len(pq.arr), InitCapacity)
	}
	if pq.arr[0] != first {
		t.Fatalf("arr[0] = %+v, want %+v", pq.arr[0], first)
	}
	assertHeapProperty(t, pq)
}

func TestPriorityQueuePushBubblesHigherPriorityUpMultipleLevels(t *testing.T) {
	pq := New[testItem]()
	input := []testItem{
		{priority: 1, id: 1},
		{priority: 2, id: 2},
		{priority: 3, id: 3},
		{priority: 4, id: 4},
		{priority: 5, id: 5},
	}

	pushItems(pq, input...)

	expected := []testItem{
		{priority: 5, id: 5},
		{priority: 4, id: 4},
		{priority: 2, id: 2},
		{priority: 1, id: 1},
		{priority: 3, id: 3},
	}
	assertHeapItems(t, pq, expected)
	assertHeapProperty(t, pq)
}

func TestPriorityQueuePushDoesNotBubbleEqualPriorityAboveParent(t *testing.T) {
	pq := New[testItem]()
	first := testItem{priority: 10, id: 1}
	lower := testItem{priority: 5, id: 2}
	equal := testItem{priority: 10, id: 3}

	pushItems(pq, first, lower, equal)

	expected := []testItem{first, lower, equal}
	assertHeapItems(t, pq, expected)
	assertHeapProperty(t, pq)
}

func TestPriorityQueuePushHandlesZeroAndNegativePriorities(t *testing.T) {
	pq := New[testItem]()
	input := []testItem{
		{priority: -10, id: 1},
		{priority: 0, id: 2},
		{priority: -5, id: 3},
		{priority: -1, id: 4},
	}

	pushItems(pq, input...)

	if pq.arr[0] != (testItem{priority: 0, id: 2}) {
		t.Fatalf("root = %+v, want priority 0 item", pq.arr[0])
	}
	assertContainsItems(t, pq, input)
	assertHeapProperty(t, pq)
}

func TestPriorityQueuePushGrowsWhenCapacityIsReached(t *testing.T) {
	pq := New[testItem]()
	pushed := make([]testItem, 0, InitCapacity+1)

	for i := 0; i < InitCapacity; i++ {
		item := testItem{priority: i, id: i}
		pushed = append(pushed, item)
		pq.Push(item)
	}

	if pq.size != InitCapacity {
		t.Fatalf("size before growth = %d, want %d", pq.size, InitCapacity)
	}
	if pq.capacity != InitCapacity {
		t.Fatalf("capacity before growth = %d, want %d", pq.capacity, InitCapacity)
	}
	if len(pq.arr) != InitCapacity {
		t.Fatalf("len(arr) before growth = %d, want %d", len(pq.arr), InitCapacity)
	}

	highest := testItem{priority: InitCapacity + 1, id: InitCapacity + 1}
	pushed = append(pushed, highest)
	pq.Push(highest)

	if pq.size != InitCapacity+1 {
		t.Fatalf("size after growth = %d, want %d", pq.size, InitCapacity+1)
	}
	if pq.capacity != InitCapacity*2 {
		t.Fatalf("capacity after growth = %d, want %d", pq.capacity, InitCapacity*2)
	}
	if len(pq.arr) != InitCapacity*2 {
		t.Fatalf("len(arr) after growth = %d, want %d", len(pq.arr), InitCapacity*2)
	}
	if pq.arr[0] != highest {
		t.Fatalf("root after growth = %+v, want %+v", pq.arr[0], highest)
	}
	assertContainsItems(t, pq, pushed)
	assertHeapProperty(t, pq)
}

func TestPriorityQueueIsEmptyReportsQueueState(t *testing.T) {
	pq := New[testItem]()

	if !pq.IsEmpty() {
		t.Fatal("IsEmpty() = false, want true for new queue")
	}

	pq.Push(testItem{priority: 1, id: 1})

	if pq.IsEmpty() {
		t.Fatal("IsEmpty() = true, want false after Push")
	}
}

func TestPriorityQueueLenReportsNumberOfItems(t *testing.T) {
	pq := New[testItem]()

	if pq.Len() != 0 {
		t.Fatalf("Len() = %d, want 0 for new queue", pq.Len())
	}

	input := []testItem{
		{priority: 1, id: 1},
		{priority: 3, id: 2},
		{priority: 2, id: 3},
	}
	for i, item := range input {
		pq.Push(item)
		if pq.Len() != i+1 {
			t.Fatalf("Len() after %d pushes = %d, want %d", i+1, pq.Len(), i+1)
		}
	}
}

func TestPriorityQueuePeekReturnsFalseWhenEmpty(t *testing.T) {
	pq := New[testItem]()

	got, ok := pq.Peek()

	if ok {
		t.Fatal("Peek() ok = true, want false for empty queue")
	}
	if got != (testItem{}) {
		t.Fatalf("Peek() item = %+v, want zero value", got)
	}
}

func TestPriorityQueuePeekReturnsHighestPriorityWithoutRemoving(t *testing.T) {
	pq := New[testItem]()
	highest := testItem{priority: 10, id: 2}
	pushItems(
		pq,
		testItem{priority: 1, id: 1},
		highest,
		testItem{priority: 5, id: 3},
	)

	got, ok := pq.Peek()

	if !ok {
		t.Fatal("Peek() ok = false, want true")
	}
	if got != highest {
		t.Fatalf("Peek() item = %+v, want %+v", got, highest)
	}
	if pq.Len() != 3 {
		t.Fatalf("Len() after Peek() = %d, want 3", pq.Len())
	}
	if pq.IsEmpty() {
		t.Fatal("IsEmpty() after Peek() = true, want false")
	}
	assertHeapProperty(t, pq)
}

func TestPriorityQueuePopReturnsFalseWhenEmpty(t *testing.T) {
	pq := New[testItem]()

	got, ok := pq.Pop()

	if ok {
		t.Fatal("Pop() ok = true, want false for empty queue")
	}
	if got != (testItem{}) {
		t.Fatalf("Pop() item = %+v, want zero value", got)
	}
	if pq.Len() != 0 {
		t.Fatalf("Len() after empty Pop() = %d, want 0", pq.Len())
	}
	if !pq.IsEmpty() {
		t.Fatal("IsEmpty() after empty Pop() = false, want true")
	}
}

func TestPriorityQueuePopRemovesSingleItem(t *testing.T) {
	pq := New[testItem]()
	item := testItem{priority: 42, id: 1}
	pq.Push(item)

	got, ok := pq.Pop()

	if !ok {
		t.Fatal("Pop() ok = false, want true")
	}
	if got != item {
		t.Fatalf("Pop() item = %+v, want %+v", got, item)
	}
	if pq.Len() != 0 {
		t.Fatalf("Len() after Pop() = %d, want 0", pq.Len())
	}
	if !pq.IsEmpty() {
		t.Fatal("IsEmpty() after Pop() = false, want true")
	}
}

func TestPriorityQueuePopReturnsItemsByDescendingPriority(t *testing.T) {
	pq := New[testItem]()
	input := []testItem{
		{priority: 3, id: 1},
		{priority: 10, id: 2},
		{priority: -1, id: 3},
		{priority: 7, id: 4},
		{priority: 0, id: 5},
	}
	expected := []testItem{
		{priority: 10, id: 2},
		{priority: 7, id: 4},
		{priority: 3, id: 1},
		{priority: 0, id: 5},
		{priority: -1, id: 3},
	}

	pushItems(pq, input...)

	for i, want := range expected {
		got, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop() #%d ok = false, want true", i+1)
		}
		if got != want {
			t.Fatalf("Pop() #%d item = %+v, want %+v", i+1, got, want)
		}
		if pq.Len() != len(expected)-i-1 {
			t.Fatalf("Len() after Pop() #%d = %d, want %d", i+1, pq.Len(), len(expected)-i-1)
		}
		assertHeapProperty(t, pq)
	}

	if !pq.IsEmpty() {
		t.Fatal("IsEmpty() after popping all items = false, want true")
	}
}

func pushItems(pq *PriorityQueue[testItem], items ...testItem) {
	for _, item := range items {
		pq.Push(item)
	}
}

func assertHeapItems(t *testing.T, pq *PriorityQueue[testItem], expected []testItem) {
	t.Helper()
	if pq.size != len(expected) {
		t.Fatalf("size = %d, want %d", pq.size, len(expected))
	}
	for i, item := range expected {
		if pq.arr[i] != item {
			t.Fatalf("arr[%d] = %+v, want %+v; heap = %+v", i, pq.arr[i], item, pq.arr[:pq.size])
		}
	}
}

func assertContainsItems(t *testing.T, pq *PriorityQueue[testItem], expected []testItem) {
	t.Helper()
	if pq.size != len(expected) {
		t.Fatalf("size = %d, want %d", pq.size, len(expected))
	}

	counts := make(map[testItem]int, len(expected))
	for _, item := range expected {
		counts[item]++
	}
	for _, item := range pq.arr[:pq.size] {
		counts[item]--
	}
	for item, count := range counts {
		if count != 0 {
			t.Fatalf("item %+v count delta = %d; heap = %+v", item, count, pq.arr[:pq.size])
		}
	}
}

func assertHeapProperty(t *testing.T, pq *PriorityQueue[testItem]) {
	t.Helper()
	for childIdx := 1; childIdx < pq.size; childIdx++ {
		parentIdx := getParentIdx(childIdx)
		if pq.arr[childIdx].CompareTo(pq.arr[parentIdx]) > 0 {
			t.Fatalf(
				"heap property violated: child arr[%d]=%+v has higher priority than parent arr[%d]=%+v; heap = %+v",
				childIdx,
				pq.arr[childIdx],
				parentIdx,
				pq.arr[parentIdx],
				pq.arr[:pq.size],
			)
		}
	}
}
