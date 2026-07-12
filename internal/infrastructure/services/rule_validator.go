package services

import (
	"password-generator/internal/domain"
)

type RuleValidator struct {
	MinLength        int
	MaxLength        int
	MinSpecialChars  int
	MaxConsecutive   int
}

func NewRuleValidator() *RuleValidator {
	return &RuleValidator{
		MinLength:       8,
		MaxLength:       64,
		MinSpecialChars: 1,
		MaxConsecutive:  2,
	}
}

func (v *RuleValidator) Validate(password string) (*domain.ValidationResult, error) {
	result := domain.NewValidationResult()
	p := domain.NewPassword(password)

	v.validateLength(p, result)
	v.validateSpecialChars(p, result)
	v.validateConsecutive(p, result)

	alphabetSize := p.AlphabetSize()
	entropy := domain.CalculateEntropy(p.Length, alphabetSize)
	result.Entropy = entropy
	result.Strength = domain.ClassifyStrength(entropy)

	return result, nil
}

func (v *RuleValidator) validateLength(p *domain.Password, result *domain.ValidationResult) {
	if p.Length < v.MinLength {
		result.AddError("password must be at least " + itoa(v.MinLength) + " characters long")
	}
	if p.Length > v.MaxLength {
		result.AddError("password must be at most " + itoa(v.MaxLength) + " characters long")
	}
}

func (v *RuleValidator) validateSpecialChars(p *domain.Password, result *domain.ValidationResult) {
	if !p.HasSpecialChars() {
		result.AddError("password must contain at least one special character")
	}
}

func (v *RuleValidator) validateConsecutive(p *domain.Password, result *domain.ValidationResult) {
	count := 1
	for i := 1; i < p.Length; i++ {
		if p.Value[i] == p.Value[i-1] {
			count++
			if count > v.MaxConsecutive {
				result.AddError("password must not have more than " + itoa(v.MaxConsecutive) + " consecutive identical characters")
				return
			}
		} else {
			count = 1
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}