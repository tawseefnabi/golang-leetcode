package lengthoflastword

// Testing in go
import (
	"testing"
)

func Test_LengthOfLastWord(t *testing.T) {
	s := "Geeks For Geeks "
	expected := 5
	if res := LengthOfLastWord(s); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
	str := "   fly me   to   the moon  "
	expected1 := 4
	if res := LengthOfLastWord(str); res != expected1 {
		t.Errorf("expected %d, got %d", expected1, res)
	}
	str1 := "luffy is still joyboy"
	expected2 := 6
	if res := LengthOfLastWord(str1); res != expected2 {
		t.Errorf("expected %d, got %d", expected2, res)
	}

}
