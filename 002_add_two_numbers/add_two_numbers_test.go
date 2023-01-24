package addtwonumbers

import (
	// "fmt"
	"reflect"
	// "container/list"
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
func TestAddTwoNumbers(t *testing.T) {

	arg1 := createSingleLinkedList([]int{2, 4, 3})
	arg2 := createSingleLinkedList([]int{5, 6, 4})
	expected := createSingleLinkedList([]int{7, 0, 8})
	result := AddTwoNumbers(arg1, arg2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Failed, two sum")
	}
}
