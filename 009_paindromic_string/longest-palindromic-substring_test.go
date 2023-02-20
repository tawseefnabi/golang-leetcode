package paindromicstring

// Testing in go
import (
	"reflect"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	s1 := "babad"
	if !reflect.DeepEqual(longestPalindrome(s1), "bab") {
		t.Errorf("expected %s, got %s", "bab", longestPalindrome(s1))
	}
	s2 := "cbbd"
	if !reflect.DeepEqual(longestPalindrome(s2), "bb") {
		t.Errorf("expected %s, got %s", "bb", longestPalindrome(s2))
	}
}
