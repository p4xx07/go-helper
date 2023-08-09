package generic_helper

import (
	"testing"
	"time"
)

func Test_String_To_String(t *testing.T) {

	out, err := ConvertFromString[string]("string")

	if err != nil {
		panic(err)
	}

	if *out != "string" {
		panic("different")
	}
}

func TestConvertStringToInt(t *testing.T) {
	out, err := ConvertFromString[int]("42")
	if err != nil {
		t.Fatal(err)
	}
	if *out != 42 {
		t.Fatal("Expected: 42, Got:", out)
	}
}

func TestConvertStringToUint(t *testing.T) {
	out, err := ConvertFromString[uint]("42")
	if err != nil {
		t.Fatal(err)
	}
	if *out != 42 {
		t.Fatal("Expected: 42, Got:", out)
	}
}

func TestConvertStringToFloat(t *testing.T) {
	out, err := ConvertFromString[float32]("3.14")
	if err != nil {
		t.Fatal(err)
	}
	if *out != 3.14 {
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
		out, err := ConvertFromString[time.Time](test.input)
		if err != nil {
			t.Fatalf("Error converting %s: %v", test.input, err)
		}

		if !out.Equal(test.expected) {
			t.Errorf("Conversion mismatch for %s. Expected: %v, Got: %v", test.input, test.expected, out)
		}
	}
}

func TestConvertStringToUnsupportedType(t *testing.T) {
	out, err := ConvertFromString[complex128]("invalid")
	if out != nil {
		t.Fatal("should be nil")
	}
	if err == nil {
		t.Fatal("Expected error for unsupported type, but got nil")
	}
}
