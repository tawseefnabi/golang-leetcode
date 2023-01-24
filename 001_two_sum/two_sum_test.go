package two_sum

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	res := TwoSum(nums, 9)
	if !reflect.DeepEqual(res, []int{0, 2}) {
		t.Error("Failed, two sum")
	}
	nums2 := []int{3, 2, 4}
	res2 := TwoSum(nums2, 6)
	if !reflect.DeepEqual(res2, []int{1, 2}) {
		t.Error("Failed, two sum")
	}
	nums3 := []int{3, 3}
	res3 := TwoSum(nums3, 6)
	if !reflect.DeepEqual(res3, []int{0, 1}) {
		t.Error("Failed, two sum")
	}
}

func TestTwoSum_new(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	res := TwoSum_new(nums, 9)
	if !reflect.DeepEqual(res, []int{0, 2}) {
		t.Error("Failed, two sum")
	}
	nums2 := []int{3, 2, 4}
	res2 := TwoSum_new(nums2, 6)
	if !reflect.DeepEqual(res2, []int{1, 2}) {
		t.Error("Failed, two sum")
	}
	nums3 := []int{3, 3}
	res3 := TwoSum_new(nums3, 6)
	if !reflect.DeepEqual(res3, []int{0, 1}) {
		t.Error("Failed, two sum")
	}
}
