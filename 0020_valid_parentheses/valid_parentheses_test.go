package validparentheses

// Testing in go
import (
	"testing"
)

func Test_isValidParentheses(t *testing.T) {
	str := "()"
	// isValidParentheses(str)
	if res := isValidParentheses(str); !res {
		t.Errorf("expected %t, got %t", true, res)
	}
	if res := isValidParentheses("()[]{}"); !res {
		t.Errorf("expected %t, got %t", true, res)
	}
	if res := isValidParentheses("(]"); res {
		t.Errorf("expected %t, got %t", false, res)
	}
	if res := isValidParentheses("([)]"); res {
		t.Errorf("expected %t, got %t", false, res)
	}
	if res := isValidParentheses("{[]}"); !res {
		t.Errorf("expected %t, got %t", true, res)
	}
}
