package random_helper

import (
	"testing"
	"unicode"
)

func Test_GenerateLength(t *testing.T) {
	expectedSlice := []int{
		1, 2, 3, 4, 5, 10, 20, 100,
	}

	for _, expected := range expectedSlice {
		actual := Generate(expected, Numbers)
		if len(actual) != expected {
			panic("not of the same length")
		}
	}
}

func TestGenerate(t *testing.T) {
	// Test case for Complexity AZ
	result := Generate(10, AZ)
	for _, c := range result {
		if !unicode.IsLower(c) {
			t.Errorf("Generated string contains non-lowercase characters")
		}
	}

	// Test case for Complexity AZCaps
	result = Generate(10, AZCaps)
	for _, c := range result {
		if !unicode.IsUpper(c) {
			t.Errorf("Generated string contains non-uppercase characters")
		}
	}

	// Test case for Complexity AZAndCaps
	result = Generate(10, AZAndCaps)
	for _, c := range result {
		if !unicode.IsLetter(c) {
			t.Errorf("Generated string contains non-letter characters")
		}
	}

	// Test case for Complexity Numbers
	result = Generate(10, Numbers)
	for _, c := range result {
		if !unicode.IsNumber(c) {
			t.Errorf("Generated string contains non-number characters")
		}
	}

	// Test case for Complexity AZAndCapsAndNumbers
	result = Generate(10, AZAndCapsAndNumbers)
	for _, c := range result {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			t.Errorf("Generated string contains characters that are neither letters nor numbers")
		}
	}
	// Test case for Complexity AZAndCapsAndNumbersSymbols
	result = Generate(10, AZAndCapsAndNumbersSymbols)
	for _, c := range result {
		if !unicode.IsPrint(c) {
			t.Errorf("Generated string contains characters that are not printable")
		}
	}
}
