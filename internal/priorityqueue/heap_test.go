package priorityqueue

import (
	"reflect"
	"testing"
)

func intCompare(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func TestHeapInsertion(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Ascending array", []int{1, 2, 3, 4, 5}, []int{5, 4, 2, 1, 3}},
		{"Descending array", []int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{"Empty array", []int{}, []int{}},
		{"Singleton array", []int{1}, []int{1}},
		{"Random array", []int{47, 12, 58, 33, 21, 89, 6, 72, 15, 64}, []int{89, 72, 58, 33, 64, 47, 6, 12, 15, 21}},
	}

	for _, testCase := range testCases {
		h := newHeap(intCompare)
		for _, item := range testCase.input {
			h.insert(item)
		}
		if !checkEquals(testCase.expected, h) {
			t.Errorf("failed: %s\n%v\n%v", testCase.name, h.arr, testCase.expected)
		}
	}
}

func TestHeapRemoval(t *testing.T) {
	input := []int{2, 3, 5, 1, 4}
	expected := []int{5, 4, 3, 2, 1}
	h := newHeap(intCompare)
	for _, a := range input {
		h.insert(a)
	}
	output := make([]int, len(input))
	i := 0
	for {
		a, exists := h.remove()
		if !exists {
			break
		}
		output[i] = a
		i = i + 1
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Failed\n%v\n%v\n", output, expected)
	}
}

func checkEquals(arr []int, h *heap[int]) bool {
	if len(arr) != h.size {
		return false
	}
	for i, a := range arr {
		if a != h.arr[i] {
			return false
		}
	}
	return true
}
