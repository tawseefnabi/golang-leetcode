package climbingstairs

// Testing in go
import (
	// "reflect"
	// "fmt"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	if steps := climbStairs(0); steps != 1 {
		t.Errorf("expected %d, got %d", 1, steps)
	}
	if steps := climbStairs(1); steps != 1 {
		t.Errorf("expected %d, got %d", 1, steps)
	}
	if steps := climbStairs(2); steps != 2 {
		t.Errorf("expected %d, got %d", 2, steps)
	}
	if steps := climbStairs(3); steps != 3 {
		t.Errorf("expected %d, got %d", 3, steps)
	}
	if steps := climbStairs(34); steps != 9227465 {
		t.Errorf("expected %d, got %d", 9227465, steps)
	}
}
