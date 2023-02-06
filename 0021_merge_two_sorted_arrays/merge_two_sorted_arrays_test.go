package mergetwosortedarrays

// Testing in go
import (
	// "fmt"
	"reflect"
	"testing"
)

func createSingleLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, j := range arr {
		cur.Next = &ListNode{Val: j}
		cur = cur.Next
	}
	return head.Next
}
func Test_MergeTwoArrays(t *testing.T) {
	arg1 := createSingleLinkedList([]int{1, 3, 5})
	arg2 := createSingleLinkedList([]int{2, 4, 6})
	expected := createSingleLinkedList([]int{1, 2, 3, 4, 5, 6})
	if res := MergeTwoArrays(arg1, arg2); !reflect.DeepEqual(res, expected) {
		t.Errorf("mergeTwoLists() = %v, expected %v", res, expected)
	}
	_arg1 := createSingleLinkedList([]int{})
	_arg2 := createSingleLinkedList([]int{})
	_expected := createSingleLinkedList([]int{})
	if res := MergeTwoArrays(_arg1, _arg2); !reflect.DeepEqual(res, _expected) {
		t.Errorf("mergeTwoLists() = %v, expected %v", res, _expected)
	}
}
