package sqrtx

// Testing in go
import (
	"testing"
)

func Test_Sqrt(t *testing.T) {
	if res := Sqrt(9); res != 3 {
		t.Errorf("expected %d, got %d", 3, res)
	}
	if res := Sqrt(1); res != 1 {
		t.Errorf("expected %d, got %d", 1, res)
	}
	if res := Sqrt(16); res != 4 {
		t.Errorf("expected %d, got %d", 4, res)
	}
	if res := Sqrt(8); res != 2 {
		t.Errorf("expected %d, got %d", 2, res)
	}
}
