package usecase

import (
	"errors"
	"password-generator/internal/domain"
	"password-generator/internal/interfaces"
	"testing"
)

type mockGenerator struct {
	result string
	err    error
}

func (m *mockGenerator) Generate(config interfaces.GenerateConfig) (string, error) {
	return m.result, m.err
}

type mockValidator struct {
	result *domain.ValidationResult
	err    error
}

func (m *mockValidator) Validate(password string) (*domain.ValidationResult, error) {
	return m.result, m.err
}

func TestGeneratePasswordUseCase_Execute_Success(t *testing.T) {
	gen := &mockGenerator{result: "Abc123!@#"}
	val := &mockValidator{result: &domain.ValidationResult{IsValid: true}}
	uc := NewGeneratePasswordUseCase(gen, val)

	config := interfaces.GenerateConfig{
		Length:       12,
		UseUppercase: true,
		UseLowercase: true,
		UseNumbers:   true,
		UseSymbols:   true,
	}

	password, validation, err := uc.Execute(config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if password == nil {
		t.Fatal("expected password to be non-nil")
	}
	if validation == nil {
		t.Fatal("expected validation to be non-nil")
	}
}

func TestGeneratePasswordUseCase_Execute_ShortLength(t *testing.T) {
	gen := &mockGenerator{}
	val := &mockValidator{}
	uc := NewGeneratePasswordUseCase(gen, val)

	config := interfaces.GenerateConfig{Length: 4}
	_, _, err := uc.Execute(config)
	if err == nil {
		t.Fatal("expected error for short length")
	}
}

func TestGeneratePasswordUseCase_Execute_LongLength(t *testing.T) {
	gen := &mockGenerator{}
	val := &mockValidator{}
	uc := NewGeneratePasswordUseCase(gen, val)

	config := interfaces.GenerateConfig{Length: 100}
	_, _, err := uc.Execute(config)
	if err == nil {
		t.Fatal("expected error for long length")
	}
}

func TestGeneratePasswordUseCase_Execute_NoCharTypes(t *testing.T) {
	gen := &mockGenerator{}
	val := &mockValidator{}
	uc := NewGeneratePasswordUseCase(gen, val)

	config := interfaces.GenerateConfig{Length: 12}
	_, _, err := uc.Execute(config)
	if err == nil {
		t.Fatal("expected error when no character types enabled")
	}
}

func TestGeneratePasswordUseCase_Execute_GeneratorError(t *testing.T) {
	gen := &mockGenerator{err: errors.New("generator error")}
	val := &mockValidator{}
	uc := NewGeneratePasswordUseCase(gen, val)

	config := interfaces.GenerateConfig{
		Length:       12,
		UseUppercase: true,
	}
	_, _, err := uc.Execute(config)
	if err == nil {
		t.Fatal("expected error from generator")
	}
}

func TestValidatePasswordUseCase_Execute_Success(t *testing.T) {
	val := &mockValidator{result: &domain.ValidationResult{IsValid: true}}
	uc := NewValidatePasswordUseCase(val)

	result, err := uc.Execute("Abc123!@#")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result to be non-nil")
	}
}

func TestValidatePasswordUseCase_Execute_EmptyPassword(t *testing.T) {
	val := &mockValidator{}
	uc := NewValidatePasswordUseCase(val)

	_, err := uc.Execute("")
	if err == nil {
		t.Fatal("expected error for empty password")
	}
}

func TestValidatePasswordUseCase_Execute_ValidatorError(t *testing.T) {
	val := &mockValidator{err: errors.New("validator error")}
	uc := NewValidatePasswordUseCase(val)

	_, err := uc.Execute("Abc123!@#")
	if err == nil {
		t.Fatal("expected error from validator")
	}
}