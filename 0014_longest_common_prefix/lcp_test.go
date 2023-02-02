package longestcommonprefix

// Testing in go
import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	strs1 := []string{}
	strs2 := []string{"dog", "racecar", "car"}
	strs3 := []string{""}
	expected := []string{
		"fl",
		"",
		"",
		"",
	}
	if LongestCommonPrefix(strs) != expected[0] {
		t.Errorf("expected %s, got %s", expected[0], LongestCommonPrefix(strs))
	}
	if LongestCommonPrefix(strs1) != expected[1] {
		t.Errorf("expected %s, got %s", expected[1], LongestCommonPrefix(strs1))
	}
	if LongestCommonPrefix(strs2) != expected[2] {
		t.Errorf("expected %s, got %s", expected[2], LongestCommonPrefix(strs2))
	}
	if LongestCommonPrefix(strs3) != expected[3] {
		t.Errorf("expected %s, got %s", expected[3], LongestCommonPrefix(strs3))
	}
}
