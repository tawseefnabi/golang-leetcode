package msa

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 7, 0, 0, 0}
	m := 4
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6, 7}
	// msa(nums1, m, nums2, n)
	if msa(nums1, m, nums2, n); !reflect.DeepEqual(expected, nums1) {
		t.Errorf("expected %v, got %v", expected, nums1)
	}
}
