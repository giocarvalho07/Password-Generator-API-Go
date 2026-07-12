package services

import (
	"password-generator/internal/interfaces"
	"strings"
	"testing"
)

func TestCryptoGenerator_Generate_AllTypes(t *testing.T) {
	gen := NewCryptoGenerator()
	config := interfaces.GenerateConfig{
		Length:       16,
		UseUppercase: true,
		UseLowercase: true,
		UseNumbers:   true,
		UseSymbols:   true,
	}

	result, err := gen.Generate(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 16 {
		t.Errorf("expected length 16, got %d", len(result))
	}
}

func TestCryptoGenerator_Generate_UppercaseOnly(t *testing.T) {
	gen := NewCryptoGenerator()
	config := interfaces.GenerateConfig{
		Length:       10,
		UseUppercase: true,
	}

	result, err := gen.Generate(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 10 {
		t.Errorf("expected length 10, got %d", len(result))
	}
	for _, c := range result {
		if c < 'A' || c > 'Z' {
			t.Errorf("expected uppercase letter, got '%c'", c)
		}
	}
}

func TestCryptoGenerator_Generate_LowercaseOnly(t *testing.T) {
	gen := NewCryptoGenerator()
	config := interfaces.GenerateConfig{
		Length:       10,
		UseLowercase: true,
	}

	result, err := gen.Generate(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, c := range result {
		if c < 'a' || c > 'z' {
			t.Errorf("expected lowercase letter, got '%c'", c)
		}
	}
}

func TestCryptoGenerator_Generate_NumbersOnly(t *testing.T) {
	gen := NewCryptoGenerator()
	config := interfaces.GenerateConfig{
		Length:     10,
		UseNumbers: true,
	}

	result, err := gen.Generate(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, c := range result {
		if c < '0' || c > '9' {
			t.Errorf("expected number, got '%c'", c)
		}
	}
}

func TestCryptoGenerator_Generate_SymbolsOnly(t *testing.T) {
	gen := NewCryptoGenerator()
	config := interfaces.GenerateConfig{
		Length:     10,
		UseSymbols: true,
	}

	result, err := gen.Generate(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	symbols := "!@#$%^&*()_+-=[]{}|;:',.<>?/`~"
	for _, c := range result {
		if !strings.ContainsRune(symbols, c) {
			t.Errorf("expected symbol, got '%c'", c)
		}
	}
}

func TestCryptoGenerator_Generate_NoTypes(t *testing.T) {
	gen := NewCryptoGenerator()
	config := interfaces.GenerateConfig{Length: 10}

	_, err := gen.Generate(config)
	if err == nil {
		t.Fatal("expected error when no character types enabled")
	}
}

func TestRuleValidator_Validate_Valid(t *testing.T) {
	v := NewRuleValidator()
	result, err := v.Validate("Abc123!@#")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsValid {
		t.Errorf("expected valid password, got errors: %v", result.Errors)
	}
	if result.Strength != "medium" {
		t.Errorf("expected strength 'medium', got '%s'", result.Strength)
	}
}

func TestRuleValidator_Validate_TooShort(t *testing.T) {
	v := NewRuleValidator()
	result, err := v.Validate("Ab1!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsValid {
		t.Error("expected invalid password")
	}
}

func TestRuleValidator_Validate_NoSpecialChars(t *testing.T) {
	v := NewRuleValidator()
	result, err := v.Validate("Abc12345")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsValid {
		t.Error("expected invalid password")
	}
}

func TestRuleValidator_Validate_ConsecutiveChars(t *testing.T) {
	v := NewRuleValidator()
	result, err := v.Validate("AAAbc123!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsValid {
		t.Error("expected invalid password")
	}
}

func TestRuleValidator_Validate_EntropyCalculation(t *testing.T) {
	v := NewRuleValidator()
	result, err := v.Validate("Abc123!@#")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Entropy == 0 {
		t.Error("expected entropy to be calculated")
	}
}

func TestRuleValidator_Validate_WeakPassword(t *testing.T) {
	v := NewRuleValidator()
	result, err := v.Validate("abc!@#$%")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Strength != "medium" {
		t.Errorf("expected strength 'medium', got '%s'", result.Strength)
	}
}