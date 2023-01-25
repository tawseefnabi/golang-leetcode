package longestsubstringwithoutrepeatingcharacters

import (
	"reflect"
	"testing"
)

// testings
func TestLengthOfLargestSubstring(t *testing.T) {
	str := "abcabcbb"
	expected := 3
	res := lengthOfLargestSubstring(str)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
