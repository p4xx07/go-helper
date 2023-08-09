package generic_helper

import (
	"testing"
	"time"
)

func Test_String_To_String(t *testing.T) {

	var out string
	err := ConvertFromString[string]("string", &out)

	if err != nil {
		panic(err)
	}

	if out != "string" {
		panic("different")
	}
}

func TestConvertStringToInt(t *testing.T) {
	var out int
	err := ConvertFromString[int]("42", &out)
	if err != nil {
		t.Fatal(err)
	}
	if out != 42 {
		t.Fatal("Expected: 42, Got:", out)
	}
}

func TestConvertStringToUint(t *testing.T) {
	var out uint
	err := ConvertFromString[uint]("42", &out)
	if err != nil {
		t.Fatal(err)
	}
	if out != 42 {
		t.Fatal("Expected: 42, Got:", out)
	}
}

func TestConvertStringToFloat(t *testing.T) {
	var out float32
	err := ConvertFromString[float32]("3.14", &out)
	if err != nil {
		t.Fatal(err)
	}
	if out != 3.14 {
		t.Fatal("Expected: 3.14, Got:", out)
	}
}

func TestConvertStringToTime(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Time
	}{
		{"2023-08-09T00:00:00Z", time.Date(2023, time.August, 9, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		var out time.Time
		err := ConvertFromString[time.Time](test.input, &out)
		if err != nil {
			t.Fatalf("Error converting %s: %v", test.input, err)
		}

		if !out.Equal(test.expected) {
			t.Errorf("Conversion mismatch for %s. Expected: %v, Got: %v", test.input, test.expected, out)
		}
	}
}

func TestConvertStringToUnsupportedType(t *testing.T) {
	var out complex128
	err := ConvertFromString[complex128]("invalid", &out)
	if err == nil {
		t.Fatal("Expected error for unsupported type, but got nil")
	}
}
