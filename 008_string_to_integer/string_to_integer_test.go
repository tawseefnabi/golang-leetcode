package stringtointeger

// Testing in go
import (
	"reflect"
	"testing"
)

func TestAtoi(t *testing.T) {
	if !reflect.DeepEqual(Atoi("-42"), -42) {
		t.Errorf("expected %d, got %d", -42, Atoi("-42"))
	}
	if !reflect.DeepEqual(Atoi("42"), 42) {
		t.Errorf("expected %d, got %d", 42, Atoi("42"))
	}
	if !reflect.DeepEqual(Atoi("4193 with words"), 4193) {
		t.Errorf("expected %d, got %d", 4193, Atoi("4193 with words"))
	}
}
