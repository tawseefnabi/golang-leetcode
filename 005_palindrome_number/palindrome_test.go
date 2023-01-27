package palindromenumber

import (
	"fmt"
	// "reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	if !IsPalindrome(121) {
		t.Errorf("expected %t, got %t", true, IsPalindrome(121))
	}
	if IsPalindrome(-121) != false {
		t.Errorf("expected %t, got %t", true, IsPalindrome(-121))
	}
	if IsPalindrome(10) != false {
		t.Errorf("expected %t, got %t", true, IsPalindrome(10))
	}
	if !IsPalindrome(0) {
		t.Errorf("expected %t, got %t", true, IsPalindrome(0))
	}
}
