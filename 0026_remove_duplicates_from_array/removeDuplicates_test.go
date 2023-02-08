package removeduplicatesfromarray

import (
	"testing"
)

func Test_MergeTwoArrays(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	if res := removeDuplicates(arr); res != 5 {
		t.Errorf("expected %d, got %d", 5, res)
	}
	arr2 := []int{}
	if res := removeDuplicates(arr2); res != 0 {
		t.Errorf("expected %d, got %d", 0, res)
	}
	arr1 := []int{1, 1, 2}
	if res := removeDuplicates(arr1); res != 2 {
		t.Errorf("expected %d, got %d", 2, res)
	}
}
