package math

import (
	"testing"
)

// TestAbs is doing something
func TestAbs(t *testing.T) {
	case1 := Abs(-1)
	if case1 != 1 {
		t.Errorf("Abs(-1) = %d; want 1", case1)
	}
	case2 := Abs(0)
	if case2 != 0 {
		t.Errorf("Abs(0) = %d; want ", case2)
	}
}

// TestMod is doing something
func TestMod(t *testing.T) {
	case1 := Mod(1, 1)
	if case1 != 1 {
		t.Errorf("Mod(1, 1) = %d; want 0", case1)
	}
	case2 := Mod(-5, 2)
	if case2 != 0 {
		t.Errorf("Mod(-5, 1) = %d; want ", case2)
	}
}
