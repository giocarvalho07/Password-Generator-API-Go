package domain

import "math"

type ValidationResult struct {
	IsValid  bool     `json:"is_valid"`
	Errors   []string `json:"errors"`
	Strength string   `json:"strength"`
	Entropy  float64  `json:"entropy"`
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		IsValid:  true,
		Errors:   []string{},
		Strength: "",
		Entropy:  0,
	}
}

func (r *ValidationResult) AddError(err string) {
	r.Errors = append(r.Errors, err)
	r.IsValid = false
}

func CalculateEntropy(length, alphabetSize int) float64 {
	if alphabetSize <= 0 || length <= 0 {
		return 0
	}
	return float64(length) * math.Log2(float64(alphabetSize))
}

func ClassifyStrength(entropy float64) string {
	switch {
	case entropy < 40:
		return "weak"
	case entropy < 60:
		return "medium"
	default:
		return "strong"
	}
}