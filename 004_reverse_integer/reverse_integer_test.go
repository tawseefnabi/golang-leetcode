package reverseinteger

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Print("===", Reverse(123))
	if !reflect.DeepEqual(Reverse(-123), -321) {
		t.Errorf("expected %d, got %d", -321, Reverse(-123))
	}
	if !reflect.DeepEqual(Reverse(120), 21) {
		t.Errorf("expected %d, got %d", 21, Reverse(120))
	}
	if !reflect.DeepEqual(Reverse(1523456789), 0) {
		t.Errorf("expected %d, got %d", 0, Reverse(1523456789))
	}
}
