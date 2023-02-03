package removeelement

// Testing in go
import (
	"testing"
)

func Test_RemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	if RemoveElement(nums, val) != 5 {
		t.Errorf("expected %d, got %d", 5, RemoveElement(nums, val))
	}
}
