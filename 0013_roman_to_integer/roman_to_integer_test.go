package romantointeger

// Testing in go
import (
	"reflect"
	"testing"
)

func TestAtoi(t *testing.T) {
	if !reflect.DeepEqual(RomanToInteger("IV"), 4) {
		t.Errorf("expected %d, got %d", 4, RomanToInteger("IV"))
	}
	if !reflect.DeepEqual(RomanToInteger("III"), 3) {
		t.Errorf("expected %d, got %d", 3, RomanToInteger("III"))
	}
	if !reflect.DeepEqual(RomanToInteger("IX"), 9) {
		t.Errorf("expected %d, got %d", 9, RomanToInteger("IX"))
	}
	if !reflect.DeepEqual(RomanToInteger("LVIII"), 58) {
		t.Errorf("expected %d, got %d", 58, RomanToInteger("LVIII"))
	}
	if !reflect.DeepEqual(RomanToInteger("MCMXCIV"), 1994) {
		t.Errorf("expected %d, got %d", 1994, RomanToInteger("MCMXCIV"))
	}
	if !reflect.DeepEqual(RomanToInteger(""), 0) {
		t.Errorf("expected %d, got %d", 0, RomanToInteger(""))
	}
}
