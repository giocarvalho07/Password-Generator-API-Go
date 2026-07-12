package domain

import (
	"math"
	"testing"
)

func TestNewPassword(t *testing.T) {
	p := NewPassword("abc123")
	if p.Value != "abc123" {
		t.Errorf("expected value 'abc123', got '%s'", p.Value)
	}
	if p.Length != 6 {
		t.Errorf("expected length 6, got %d", p.Length)
	}
	if p.Generated.IsZero() {
		t.Error("expected Generated to be set")
	}
}

func TestPasswordString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "a*c"},
		{"abcdef", "a****f"},
		{"ab", "**"},
		{"a", "*"},
		{"", ""},
	}

	for _, tt := range tests {
		p := NewPassword(tt.input)
		result := p.String()
		if result != tt.expected {
			t.Errorf("String(%s) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}

func TestPasswordHasUppercase(t *testing.T) {
	if !NewPassword("Abc").HasUppercase() {
		t.Error("expected HasUppercase to return true for 'Abc'")
	}
	if NewPassword("abc").HasUppercase() {
		t.Error("expected HasUppercase to return false for 'abc'")
	}
}

func TestPasswordHasLowercase(t *testing.T) {
	if !NewPassword("abc").HasLowercase() {
		t.Error("expected HasLowercase to return true for 'abc'")
	}
	if NewPassword("ABC").HasLowercase() {
		t.Error("expected HasLowercase to return false for 'ABC'")
	}
}

func TestPasswordHasNumbers(t *testing.T) {
	if !NewPassword("abc123").HasNumbers() {
		t.Error("expected HasNumbers to return true for 'abc123'")
	}
	if NewPassword("abcdef").HasNumbers() {
		t.Error("expected HasNumbers to return false for 'abcdef'")
	}
}

func TestPasswordHasSpecialChars(t *testing.T) {
	if !NewPassword("abc!@#").HasSpecialChars() {
		t.Error("expected HasSpecialChars to return true for 'abc!@#'")
	}
	if NewPassword("abcdef").HasSpecialChars() {
		t.Error("expected HasSpecialChars to return false for 'abcdef'")
	}
}

func TestPasswordAlphabetSize(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abc", 26},
		{"ABC", 26},
		{"123", 10},
		{"!@#", 32},
		{"abc123!@#", 68},
		{"abcABC123!@#", 94},
	}

	for _, tt := range tests {
		p := NewPassword(tt.input)
		result := p.AlphabetSize()
		if result != tt.expected {
			t.Errorf("AlphabetSize(%s) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestNewValidationResult(t *testing.T) {
	r := NewValidationResult()
	if !r.IsValid {
		t.Error("expected IsValid to be true")
	}
	if len(r.Errors) != 0 {
		t.Errorf("expected 0 errors, got %d", len(r.Errors))
	}
	if r.Strength != "" {
		t.Errorf("expected empty strength, got '%s'", r.Strength)
	}
	if r.Entropy != 0 {
		t.Errorf("expected entropy 0, got %f", r.Entropy)
	}
}

func TestValidationResultAddError(t *testing.T) {
	r := NewValidationResult()
	r.AddError("test error")
	if r.IsValid {
		t.Error("expected IsValid to be false after AddError")
	}
	if len(r.Errors) != 1 {
		t.Errorf("expected 1 error, got %d", len(r.Errors))
	}
	if r.Errors[0] != "test error" {
		t.Errorf("expected error 'test error', got '%s'", r.Errors[0])
	}
}

func TestCalculateEntropy(t *testing.T) {
	tests := []struct {
		length       int
		alphabetSize int
		expected     float64
	}{
		{8, 26, 37.603518},
		{12, 62, 71.450356},
		{16, 94, 104.873422},
		{0, 26, 0},
		{8, 0, 0},
		{-1, 26, 0},
	}

	for _, tt := range tests {
		result := CalculateEntropy(tt.length, tt.alphabetSize)
		if math.Abs(result-tt.expected) > 0.0001 {
			t.Errorf("CalculateEntropy(%d, %d) = %f, want %f", tt.length, tt.alphabetSize, result, tt.expected)
		}
	}
}

func TestClassifyStrength(t *testing.T) {
	tests := []struct {
		entropy  float64
		expected string
	}{
		{0, "weak"},
		{39, "weak"},
		{40, "medium"},
		{59, "medium"},
		{60, "strong"},
		{100, "strong"},
	}

	for _, tt := range tests {
		result := ClassifyStrength(tt.entropy)
		if result != tt.expected {
			t.Errorf("ClassifyStrength(%f) = %s, want %s", tt.entropy, result, tt.expected)
		}
	}
}