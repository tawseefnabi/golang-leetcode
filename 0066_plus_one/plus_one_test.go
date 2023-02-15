package plusone

// Testing in go
import (
	// "fmt"
	"reflect"
	"testing"
)

func Test_PlusOne(t *testing.T) {
	digits := []int{9, 9, 9, 9}
	expected := []int{1, 0, 0, 0, 0}
	if res := PlusOne(digits); !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
	digits1 := []int{1, 2, 3}
	expected1 := []int{1, 2, 4}
	if res1 := PlusOne(digits1); !reflect.DeepEqual(expected1, res1) {
		t.Errorf("expected %v, got %v", expected1, res1)
	}
	digits2 := []int{9}
	expected2 := []int{1, 0}
	if res2 := PlusOne(digits2); !reflect.DeepEqual(expected2, res2) {
		t.Errorf("expected %v, got %v", expected2, res2)
	}
}
