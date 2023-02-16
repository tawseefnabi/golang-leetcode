package addbinary

import (
	// "fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	if res := addBinary("11", "1"); res != "100" {
		t.Errorf("expected %s, got %s", "100", res)
	}
	if res := addBinary("1", "11"); res != "100" {
		t.Errorf("expected %s, got %s", "100", res)
	}
	if res := addBinary("1010", "1011"); res != "10101" {
		t.Errorf("expected %s, got %s", "10101", res)
	}
}
