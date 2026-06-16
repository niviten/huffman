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

func TestPriorityQueueInsertAddsFirstItem(t *testing.T) {
	pq := New[testItem]()
	first := testItem{priority: 42, id: 1}

	pq.Insert(first)

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

func TestPriorityQueueInsertBubblesHigherPriorityUpMultipleLevels(t *testing.T) {
	pq := New[testItem]()
	input := []testItem{
		{priority: 1, id: 1},
		{priority: 2, id: 2},
		{priority: 3, id: 3},
		{priority: 4, id: 4},
		{priority: 5, id: 5},
	}

	insertItems(pq, input...)

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

func TestPriorityQueueInsertDoesNotBubbleEqualPriorityAboveParent(t *testing.T) {
	pq := New[testItem]()
	first := testItem{priority: 10, id: 1}
	lower := testItem{priority: 5, id: 2}
	equal := testItem{priority: 10, id: 3}

	insertItems(pq, first, lower, equal)

	expected := []testItem{first, lower, equal}
	assertHeapItems(t, pq, expected)
	assertHeapProperty(t, pq)
}

func TestPriorityQueueInsertHandlesZeroAndNegativePriorities(t *testing.T) {
	pq := New[testItem]()
	input := []testItem{
		{priority: -10, id: 1},
		{priority: 0, id: 2},
		{priority: -5, id: 3},
		{priority: -1, id: 4},
	}

	insertItems(pq, input...)

	if pq.arr[0] != (testItem{priority: 0, id: 2}) {
		t.Fatalf("root = %+v, want priority 0 item", pq.arr[0])
	}
	assertContainsItems(t, pq, input)
	assertHeapProperty(t, pq)
}

func TestPriorityQueueInsertGrowsWhenCapacityIsReached(t *testing.T) {
	pq := New[testItem]()
	inserted := make([]testItem, 0, InitCapacity+1)

	for i := 0; i < InitCapacity; i++ {
		item := testItem{priority: i, id: i}
		inserted = append(inserted, item)
		pq.Insert(item)
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
	inserted = append(inserted, highest)
	pq.Insert(highest)

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
	assertContainsItems(t, pq, inserted)
	assertHeapProperty(t, pq)
}

func insertItems(pq *PriorityQueue[testItem], items ...testItem) {
	for _, item := range items {
		pq.Insert(item)
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
