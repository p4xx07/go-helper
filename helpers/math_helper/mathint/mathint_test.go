package mathint

import (
	"testing"
)

func TestMin(t *testing.T) {
	if Min(3, 5) != 3 {
		t.Errorf("Min(3, 5) should return 3")
	}
	if Min(-2, 4) != -2 {
		t.Errorf("Min(-2, 4) should return -2")
	}
}

func TestMax(t *testing.T) {
	if Max(3, 5) != 5 {
		t.Errorf("Max(3, 5) should return 5")
	}
	if Max(-2, 4) != 4 {
		t.Errorf("Max(-2, 4) should return 4")
	}
}

func TestAbs(t *testing.T) {
	if Abs(3) != 3 {
		t.Errorf("Abs(3) should return 3")
	}
	if Abs(-2) != 2 {
		t.Errorf("Abs(-2) should return 2")
	}
}

func TestCeil(t *testing.T) {
	if Ceil(3) != 3 {
		t.Errorf("Ceil(3) should return 3")
	}
}

func TestFloor(t *testing.T) {
	if Floor(3) != 3 {
		t.Errorf("Floor(3) should return 3")
	}
}

func TestRound(t *testing.T) {
	if Round(3) != 3 {
		t.Errorf("Round(3) should return 3")
	}
}

func TestPow(t *testing.T) {
	if Pow(2, 3) != 8 {
		t.Errorf("Pow(2, 3) should return 8")
	}
	if Pow(3, 0) != 1 {
		t.Errorf("Pow(3, 0) should return 1")
	}
}

func TestSqrt(t *testing.T) {
	if Sqrt(4) != 2 {
		t.Errorf("Sqrt(4) should return 2")
	}
}

func TestMod(t *testing.T) {
	if Mod(5, 3) != 2 {
		t.Errorf("Mod(5, 3) should return 2")
	}
	if Mod(4, 2) != 0 {
		t.Errorf("Mod(4, 2) should return 0")
	}
}
