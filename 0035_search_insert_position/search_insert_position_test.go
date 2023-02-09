package searchinsertposition

// Testing in go
import (
	// "reflect"
	// "fmt"
	"testing"
)

func Test_SearchInsertPosition(t *testing.T) {
	testCases := [][]int{
		{1, 3, 5, 6},
		{1, 3, 5, 6},
		{1, 3, 5, 6},
		{1, 3, 5, 6},
	}
	targets := []int{5, 2, 7, 0}

	expected := []int{2, 1, 4, 0}

	for index, nums := range testCases {
		if res := SearchInsertPosition(nums, targets[index]); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}

}
